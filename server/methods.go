package server

import (
	"context"
	"fmt"

	"github.com/kwilteam/kwil-extensions/types"
)

// MethodFunc is a function that executes a method
type MethodFunc func(ctx *types.ExecutionContext, inputs ...*types.ScalarValue) ([]*types.ScalarValue, error)

// InitializeFunc is a function that creates a new instance of an extension.
// In most cases, this should just validate the metadata being sent.
type InitializeFunc func(ctx context.Context, metadata map[string]string) (map[string]string, error)

// WithInputsCheck checks the number of inputs.
// If the number of inputs is not equal to numInputs, it returns an error.
func WithInputsCheck(fn MethodFunc, numInputs int) MethodFunc {
	return func(ctx *types.ExecutionContext, inputs ...*types.ScalarValue) ([]*types.ScalarValue, error) {
		if len(inputs) != numInputs {
			return nil, fmt.Errorf("expected %d args, got %d", numInputs, len(inputs))
		}
		return fn(ctx, inputs...)
	}
}

// WithOutputsCheck checks the number of outputs.
// If the number of outputs is not equal to numOutputs, it returns an error.
func WithOutputsCheck(fn MethodFunc, numOutputs int) MethodFunc {
	return func(ctx *types.ExecutionContext, inputs ...*types.ScalarValue) ([]*types.ScalarValue, error) {
		res, err := fn(ctx, inputs...)
		if err != nil {
			return nil, err
		}

		if len(res) != numOutputs {
			return nil, fmt.Errorf("expected %d returns, got %d", numOutputs, len(res))
		}

		return res, nil
	}
}
