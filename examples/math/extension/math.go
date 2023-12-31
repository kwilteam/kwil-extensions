package extension

import (
	"context"
	"fmt"
	"math/big"

	"github.com/kwilteam/kwil-extensions/types"
)

type MathExtension struct{}

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

	bigVal1 := newBigFloat(float64(val0Int))

	bigVal2 := newBigFloat(float64(val1Int))

	result := new(big.Float).Quo(bigVal1, bigVal2)

	var IntResult *big.Int
	if ctx.Metadata["round"] == "up" {
		IntResult = roundUp(result)
	} else {
		IntResult = roundDown(result)
	}

	return encodeScalarValues(IntResult.Int64())
}

// roundUp takes a big.Float and returns a new big.Float rounded up.
func roundUp(f *big.Float) *big.Int {
	c := new(big.Float).SetPrec(precision).Copy(f)
	r := new(big.Int)
	f.Int(r)

	if c.Sub(c, new(big.Float).SetPrec(precision).SetInt(r)).Sign() > 0 {
		r.Add(r, big.NewInt(1))
	}

	return r
}

// roundDown takes a big.Float and returns a new big.Float rounded down.
func roundDown(f *big.Float) *big.Int {
	r := new(big.Int)
	f.Int(r)

	return r
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

// this initialize function checks if round is set.  If not, it sets it to "up"
func initialize(ctx context.Context, metadata map[string]string) (map[string]string, error) {
	_, ok := metadata["round"]
	if !ok {
		metadata["round"] = "up"
	}

	roundVal := metadata["round"]
	if roundVal != "up" && roundVal != "down" {
		return nil, fmt.Errorf("round must be either 'up' or 'down'. default is 'up'")
	}

	return metadata, nil
}

const (
	precision = 128
)

func newBigFloat(num float64) *big.Float {
	bg := new(big.Float).SetPrec(precision)

	return bg.SetFloat64(num)
}
