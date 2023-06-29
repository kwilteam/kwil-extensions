package hello

import (
	"fmt"

	"github.com/kwilteam/kwil-extensions/types"
)

type HelloWorldExt struct {
	greeting string
}

var requiredMetadata = map[string]string{
	"punctuation": "",
}

func (e *HelloWorldExt) Name() string {
	return "hello-world"
}

func (h *HelloWorldExt) SayHello(ctx *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 value for method SayHello, got %d", len(values))
	}

	if values[0].Type != types.ScalarType_STRING {
		return nil, fmt.Errorf("expected first value to be of type STRING, got %s", values[0].Type)
	}

	name := values[0].String()

	result := h.sayHello(name, ctx.Metadata["punctuation"])

	return encodeScalarValues(result)
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

func (h *HelloWorldExt) sayHello(name, punctuation string) string {
	return fmt.Sprintf("%s %s%s", h.greeting, name, punctuation)
}

func (h *HelloWorldExt) SayGoodbye(ctx *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 value for method SayGoodbye, got %d", len(values))
	}

	if values[0].Type != types.ScalarType_STRING {
		return nil, fmt.Errorf("expected first value to be of type STRING, got %s", values[0].Type)
	}

	name := string(values[0].Value)

	result := h.sayGoodbye(name, ctx.Metadata["punctuation"])

	return encodeScalarValues(result)
}

func (h *HelloWorldExt) sayGoodbye(name, punctuation string) string {
	return fmt.Sprintf("Goodbye %s%s", name, punctuation)
}
