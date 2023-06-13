package types

import (
	"context"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"

	"github.com/cstockton/go-conv"
)

type ScalarType string

const (
	ScalarType_STRING ScalarType = "STRING"
	ScalarType_INT    ScalarType = "INT"
)

func (s ScalarType) String() string {
	return string(s)
}

func ScalarTypeFromString(s string) (ScalarType, error) {
	switch strings.ToUpper(s) {
	case "STRING":
		return ScalarType_STRING, nil
	case "INT":
		return ScalarType_INT, nil
	default:
		return "", fmt.Errorf("invalid scalar type: %s", s)
	}
}

type ScalarValue struct {
	Type  ScalarType
	Value []byte
}

func NewScalarValue(v any) (*ScalarValue, error) {
	valueType := reflect.TypeOf(v)
	switch valueType.Kind() {
	case reflect.String, reflect.Float32, reflect.Float64:
		strVal, err := conv.String(v)
		if err != nil {
			return nil, fmt.Errorf("error converting string: %s", err.Error())
		}

		return &ScalarValue{
			Type:  ScalarType_STRING,
			Value: []byte(strVal),
		}, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		int64Val, err := conv.Int64(v)
		if err != nil {
			return nil, fmt.Errorf("error converting int: %s", err.Error())
		}

		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(int64Val))

		return &ScalarValue{
			Type:  ScalarType_INT,
			Value: b,
		}, nil
	default:
		return nil, fmt.Errorf("invalid scalar type: %s", valueType.Kind())
	}
}

// String returns the string representation of the value.
func (s *ScalarValue) String() string {
	return string(s.Value)
}

// Int returns the int representation of the value.
func (s *ScalarValue) Int() int64 {
	return int64(binary.LittleEndian.Uint64(s.Value))
}

// Any returns the value as an interface{}, which can be casted to the appropriate type.
func (s *ScalarValue) Any() any {
	switch s.Type {
	case ScalarType_STRING:
		return s.String()
	case ScalarType_INT:
		return s.Int()
	default:
		return nil
	}
}

type ExecutionContext struct {
	Ctx      context.Context
	Metadata map[string]string
}
