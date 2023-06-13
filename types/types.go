package types

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

type ScalarType string

const (
	ScalarType_NULL   ScalarType = "NULL"
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
	case reflect.String:
		return &ScalarValue{
			Type:  ScalarType_STRING,
			Value: []byte(v.(string)),
		}, nil
	case reflect.Int:
		return &ScalarValue{
			Type:  ScalarType_INT,
			Value: []byte(fmt.Sprintf("%d", v.(int))),
		}, nil
	default:
		return nil, fmt.Errorf("invalid scalar type: %s", valueType.Kind())
	}
}

type ExecutionContext struct {
	Ctx      context.Context
	Metadata map[string]string
}
