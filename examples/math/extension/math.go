package extension

import (
	"fmt"
	"math/big"

	"github.com/kwilteam/kwil-extensions/types"
)

type MathExtension struct{}

var requiredMetadata = map[string]string{
	"round": "up", // can be up or down
}

func (e *MathExtension) Name() string {
	return "math"
}

func (e *MathExtension) Add(ctx *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values for method Add, got %d", len(values))
	}

	val0Int, err := values[0].Int()
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w. \nreceived value: %v", err, val0Int)
	}

	val1Int, err := values[1].Int()
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w. \nreceived value: %v", err, val1Int)
	}

	return encodeScalarValues(val0Int + val1Int)
}

func (e *MathExtension) Subtract(ctx *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values for method Subtract, got %d", len(values))
	}

	val0Int, err := values[0].Int()
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w. \nreceived value: %v", err, val0Int)
	}

	val1Int, err := values[1].Int()
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w. \nreceived value: %v", err, val1Int)
	}

	return encodeScalarValues(val0Int - val1Int)
}

func (e *MathExtension) Multiply(ctx *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values for method Multiply, got %d", len(values))
	}

	val0Int, err := values[0].Int()
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w. \nreceived value: %v", err, val0Int)
	}

	val1Int, err := values[1].Int()
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w. \nreceived value: %v", err, val1Int)
	}

	return encodeScalarValues(val0Int * val1Int)
}

func (e *MathExtension) Divide(ctx *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values for method Divide, got %d", len(values))
	}

	val0Int, err := values[0].Int()
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w. \nreceived value: %v", err, val0Int)
	}

	val1Int, err := values[1].Int()
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w. \nreceived value: %v", err, val1Int)
	}

	bigVal1 := big.NewFloat(float64(val0Int))

	bigVal2 := big.NewFloat(float64(val1Int))

	result := new(big.Float).Quo(bigVal1, bigVal2)

	var IntResult *big.Int
	if ctx.Metadata["round"] == "up" {
		IntResult = roundUp(result)
	} else {
		IntResult = roundDown(result)
	}

	return encodeScalarValues(IntResult.Int64())
}

func roundUp(f *big.Float) *big.Int {
	half := new(big.Float).SetFloat64(0.5)
	if f.Cmp(big.NewFloat(0)) == -1 { // f < 0
		f.Sub(f, half)
	} else {
		f.Add(f, half)
	}

	rounded := new(big.Float).Quo(f, big.NewFloat(1))

	i := new(big.Int)
	rounded.Int(i) // get the integral part
	return i
}

func roundDown(f *big.Float) *big.Int {
	half := new(big.Float).SetFloat64(0.5)
	if f.Cmp(big.NewFloat(0)) == -1 { // f < 0
		f.Add(f, half)
	} else {
		f.Sub(f, half)
	}

	rounded := new(big.Float).Quo(f, big.NewFloat(1))

	i := new(big.Int)
	rounded.Int(i) // get the integral part
	return i
}

func encodeScalarValues(values ...any) ([]*types.ScalarValue, error) {
	scalarValues := make([]*types.ScalarValue, len(values))
	for i, v := range values {
		scalarValue, err := types.NewScalarValue(v)
		if err != nil {
			return nil, err
		}

		scalarValues[i] = scalarValue
	}

	return scalarValues, nil
}
