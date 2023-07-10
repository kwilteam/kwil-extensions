package extension

import (
	"log"

	"github.com/kwilteam/kwil-extensions/server"
)

func NewMathExtension(logger *log.Logger) (*server.ExtensionServer, error) {
	ext := &MathExtension{}

	return server.Builder().
		Named(ext.Name()).
		WithInitializer(initialize).
		WithMethods(
			map[string]server.MethodFunc{
				"add": ext.Add,
				"sub": ext.Subtract,
				"mul": ext.Multiply,
				"div": ext.Divide,
			},
		).
		WithLoggerFunc(func(l string) {
			logger.Printf("log received from extension: %s", l)
		}).
		Build()
}
