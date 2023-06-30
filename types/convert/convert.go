package convert

import (
	"encoding/json"

	gen "github.com/kwilteam/kwil-extensions/gen"
	"github.com/kwilteam/kwil-extensions/types"
)

func ConvertScalarToPb(vals []*types.ScalarValue) ([]*gen.ScalarValue, error) {
	convertedOutputs := make([]*gen.ScalarValue, 0, len(vals))
	for _, output := range vals {
		bts, err := json.Marshal(output.Value)
		if err != nil {
			return nil, err
		}

		convertedOutputs = append(convertedOutputs, &gen.ScalarValue{
			Value: bts,
		})
	}

	return convertedOutputs, nil
}

func ConvertScalarFromPb(vals []*gen.ScalarValue) ([]*types.ScalarValue, error) {
	convertedInputs := make([]*types.ScalarValue, 0, len(vals))
	for _, input := range vals {
		var v any
		err := json.Unmarshal(input.Value, &v)
		if err != nil {
			return nil, err
		}

		convertedInputs = append(convertedInputs, &types.ScalarValue{
			Value: v,
		})
	}

	return convertedInputs, nil
}
