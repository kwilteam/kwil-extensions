package types

import (
	"context"
	"fmt"
	"reflect"

	"github.com/cstockton/go-conv"
)

type ScalarValue struct {
	Value any
}

func NewScalarValue(v any) (*ScalarValue, error) {
	valueType := reflect.TypeOf(v)
	switch valueType.Kind() {
	case reflect.String, reflect.Float32, reflect.Float64:
		return &ScalarValue{
			Value: v,
		}, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &ScalarValue{
			Value: v,
		}, nil
	default:
		return nil, fmt.Errorf("invalid scalar type: %s", valueType.Kind())
	}
}

// String returns the string representation of the value.
func (s *ScalarValue) String() (string, error) {
	return conv.String(s.Value)
}

// Int returns the int representation of the value.
func (s *ScalarValue) Int() (int64, error) {
	return conv.Int64(s.Value)
}

type ExecutionContext struct {
	Ctx      context.Context
	Metadata map[string]string
}
