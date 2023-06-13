package types

import (
	"fmt"
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
