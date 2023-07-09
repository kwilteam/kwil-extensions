package server

import (
	"context"
	"fmt"
	"strings"

	gen "github.com/kwilteam/kwil-extensions/gen"
	"github.com/kwilteam/kwil-extensions/types"
	"github.com/kwilteam/kwil-extensions/types/convert"
)

// Extension is a grpc server that implements the Kwil extension
// protobuf interface.
type Extension struct {
	gen.UnimplementedExtensionServiceServer
	conf *ExtensionConfig
}

// ExtensionConfig configures the functionality of an extension.  This includes things like the extension name, the
// available methods, etc.
type ExtensionConfig struct {
	name             string
	requiredMetadata map[string]string
	methods          map[string]MethodFunc
}

func (s *Extension) Name(ctx context.Context, req *gen.NameRequest) (*gen.NameResponse, error) {
	return &gen.NameResponse{
		Name: s.conf.name,
	}, nil
}

func (s *Extension) ListMethods(ctx context.Context, req *gen.ListMethodsRequest) (*gen.ListMethodsResponse, error) {
	methods := []string{}
	for name := range s.conf.methods {
		methods = append(methods, name)
	}

	return &gen.ListMethodsResponse{
		Methods: methods,
	}, nil
}

func (s *Extension) Execute(ctx context.Context, req *gen.ExecuteRequest) (*gen.ExecuteResponse, error) {
	method, ok := s.conf.methods[strings.ToLower(req.Name)]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", req.Name)
	}

	var err error
	req.Metadata, err = mergeStringMaps(s.conf.requiredMetadata, req.Metadata)
	if err != nil {
		return nil, fmt.Errorf("error with provided metadata: %s", err.Error())
	}

	convertedInputs, err := convert.ConvertScalarFromPb(req.Args)
	if err != nil {
		return nil, fmt.Errorf("error with provided inputs: %s", err.Error())
	}

	outputs, err := method(&types.ExecutionContext{
		Ctx:      ctx,
		Metadata: req.Metadata,
	}, convertedInputs...)
	if err != nil {
		return nil, fmt.Errorf("error executing method: %s", err.Error())
	}

	convertedOutputs, err := convert.ConvertScalarToPb(outputs)
	if err != nil {
		return nil, fmt.Errorf("error converting outputs: %s", err.Error())
	}

	return &gen.ExecuteResponse{
		Outputs: convertedOutputs,
	}, nil
}

func (s *Extension) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	return &gen.GetMetadataResponse{
		Metadata: s.conf.requiredMetadata,
	}, nil
}

// mergeStringMaps merges two maps of strings.  If a key exists in both maps,
// the value from the second map is used.
// If a value in the first map is an empty string, the value from the second
// map is required to be non-empty.
func mergeStringMaps(firstMap map[string]string, secondMap map[string]string) (map[string]string, error) {
	finalMap := make(map[string]string)

	for key, value := range firstMap {
		secondValue, ok := secondMap[key]
		if !ok {
			if value == "" {
				return nil, fmt.Errorf("missing required value: %s", key)
			}

			finalMap[key] = value
		} else {
			finalMap[key] = secondValue
		}
	}

	return finalMap, nil
}
