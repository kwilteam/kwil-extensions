package server

import (
	"context"
	"fmt"
	"strings"

	gen "github.com/kwilteam/kwil-extensions/gen"
	"github.com/kwilteam/kwil-extensions/types"
	"github.com/kwilteam/kwil-extensions/types/convert"
)

type Server struct {
	gen.UnimplementedExtensionServiceServer

	Name             string
	Methods          map[string]MethodFunc
	RequiredMetadata map[string]string

	configured bool
}

func NewExtensionServer(ext *ExtensionConfig) (*Server, error) {
	return &Server{
		Name:             ext.Name,
		Methods:          ext.Methods,
		RequiredMetadata: ext.RequiredMetadata,
	}, nil
}

func (s *Server) ListMethods(ctx context.Context, req *gen.ListMethodsRequest) (*gen.ListMethodsResponse, error) {
	methods := []string{}
	for name := range s.Methods {
		methods = append(methods, name)
	}

	return &gen.ListMethodsResponse{
		Methods: methods,
	}, nil
}

func (s *Server) Execute(ctx context.Context, req *gen.ExecuteRequest) (*gen.ExecuteResponse, error) {
	if !s.configured {
		return nil, fmt.Errorf("extension has not been configured by node")
	}

	method, ok := s.Methods[strings.ToLower(req.Name)]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", req.Name)
	}

	var err error
	req.Metadata, err = mergeStringMaps(s.RequiredMetadata, req.Metadata)
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

func (s *Server) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	return &gen.GetMetadataResponse{
		Metadata: s.RequiredMetadata,
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
