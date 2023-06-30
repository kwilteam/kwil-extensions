package extension

import (
	"github.com/kwilteam/kwil-extensions/server"
	"github.com/kwilteam/kwil-extensions/server/builder"
)

func NewMathExtension() (*server.Server, error) {
	ext := &MathExtension{}

	return builder.Builder().
		Named(ext.Name()).
		WithRequiredMetadata(requiredMetadata).
		WithMethods(
			map[string]server.MethodFunc{
				"add": ext.Add,
				"sub": ext.Subtract,
				"mul": ext.Multiply,
				"div": ext.Divide,
			},
		).Build()
}
