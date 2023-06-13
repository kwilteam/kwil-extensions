package convert

import (
	"fmt"

	gen "github.com/kwilteam/kwil-extensions/gen"
	"github.com/kwilteam/kwil-extensions/types"
)

func ConvertScalarToPb(vals []*types.ScalarValue) ([]*gen.ScalarValue, error) {
	convertedOutputs := make([]*gen.ScalarValue, 0, len(vals))
	for _, output := range vals {
		convertedOutputs = append(convertedOutputs, &gen.ScalarValue{
			Type:  output.Type.String(),
			Value: output.Value,
		})
	}

	return convertedOutputs, nil
}

func ConvertScalarFromPb(vals []*gen.ScalarValue) ([]*types.ScalarValue, error) {
	convertedInputs := make([]*types.ScalarValue, 0, len(vals))
	for _, input := range vals {
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
