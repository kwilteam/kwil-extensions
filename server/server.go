package server

import (
	"context"
	"fmt"
	"strings"

	gen "github.com/kwilteam/kwil-extensions/gen"
	"github.com/kwilteam/kwil-extensions/types"
)

type Server struct {
	gen.UnimplementedExtensionServiceServer

	Extension        UserExtension
	Configs          map[string]string
	Methods          map[string]*Method
	RequiredMetadata map[string]string

	configured bool
}

func NewExtensionServer(ext *ExtensionConfig, extensionObject UserExtension) *Server {
	mappedMethods := make(map[string]*Method)
	for _, method := range ext.Methods {
		mappedMethods[strings.ToLower(method.Name)] = method
	}

	return &Server{
		Extension:        extensionObject,
		Configs:          ext.Configs,
		Methods:          mappedMethods,
		RequiredMetadata: ext.RequiredMetadata,
	}
}

func (s *Server) Configure(ctx context.Context, req *gen.ConfigureRequest) (*gen.ConfigureResponse, error) {
	mergedConfigs, err := mergeStringMaps(s.Configs, req.Config)
	if err != nil {
		return &gen.ConfigureResponse{
			Success: false,
		}, fmt.Errorf("error with provided config: %s", err.Error())
	}

	err = s.Extension.Configure(mergedConfigs)
	if err != nil {
		return &gen.ConfigureResponse{
			Success: false,
		}, fmt.Errorf("error configuring extension: %s", err.Error())
	}

	s.Configs = mergedConfigs

	s.configured = true

	return &gen.ConfigureResponse{
		Success: true,
	}, nil
}

func (s *Server) ListMethods(ctx context.Context, req *gen.ListMethodsRequest) (*gen.ListMethodsResponse, error) {
	methods := make([]*gen.Method, 0, len(s.Methods))
	for _, method := range s.Methods {
		methods = append(methods, &gen.Method{
			Name:           method.Name,
			RequiredInputs: convertRequiredInputs(method.RequiredInputs),
		})
	}

	return &gen.ListMethodsResponse{
		Methods: methods,
	}, nil
}

func convertRequiredInputs(inputs []types.ScalarType) []string {
	requiredInputs := make([]string, 0, len(inputs))
	for _, input := range inputs {
		requiredInputs = append(requiredInputs, input.String())
	}

	return requiredInputs
}

func (s *Server) Execute(ctx context.Context, req *gen.ExecuteRequest) (*gen.ExecuteResponse, error) {
	if !s.configured {
		return nil, fmt.Errorf("extension has not been configured by node")
	}

	method, ok := s.Methods[strings.ToLower(req.Name)]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", req.Name)
	}

	if len(req.Args) != len(method.RequiredInputs) {
		return nil, fmt.Errorf("incorrect number of inputs provided: expected %d, got %d", len(method.RequiredInputs), len(req.Args))
	}

	var err error
	req.Metadata, err = mergeStringMaps(s.RequiredMetadata, req.Metadata)
	if err != nil {
		return nil, fmt.Errorf("error with provided metadata: %s", err.Error())
	}

	convertedInputs, err := convertInputsFromPb(req.Args)
	if err != nil {
		return nil, fmt.Errorf("error with provided inputs: %s", err.Error())
	}

	err = method.canExecute(convertedInputs)
	if err != nil {
		return nil, fmt.Errorf("error with provided inputs: %s", err.Error())
	}

	outputs, err := method.Function(convertedInputs, req.Metadata)
	if err != nil {
		return nil, fmt.Errorf("error executing method: %s", err.Error())
	}

	convertedOutputs, err := convertOutputsToPb(outputs)
	if err != nil {
		return nil, fmt.Errorf("error converting outputs: %s", err.Error())
	}

	return &gen.ExecuteResponse{
		Outputs: convertedOutputs,
	}, nil
}

func convertOutputsToPb(outputs []*types.ScalarValue) ([]*gen.ScalarValue, error) {
	convertedOutputs := make([]*gen.ScalarValue, 0, len(outputs))
	for _, output := range outputs {
		convertedOutputs = append(convertedOutputs, &gen.ScalarValue{
			Type:  output.Type.String(),
			Value: output.Value,
		})
	}

	return convertedOutputs, nil
}

func convertInputsFromPb(inputs []*gen.ScalarValue) ([]*types.ScalarValue, error) {
	convertedInputs := make([]*types.ScalarValue, 0, len(inputs))
	for _, input := range inputs {
		convertedType, err := types.ScalarTypeFromString(input.Type)
		if err != nil {
			return nil, fmt.Errorf("invalid scalar type: %s", err.Error())
		}

		convertedInputs = append(convertedInputs, &types.ScalarValue{
			Type:  convertedType,
			Value: input.Value,
		})
	}

	return convertedInputs, nil
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

func (s *Server) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	return &gen.GetMetadataResponse{}, nil
}
