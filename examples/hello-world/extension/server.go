package hello

import (
	"github.com/kwilteam/kwil-extensions/server"
	"github.com/kwilteam/kwil-extensions/server/builder"
)

func NewHelloWorldExtension() (*server.Server, error) {
	ext := &HelloWorldExt{}

	return builder.Builder().
		Named(ext.Name()).
		WithRequiredMetadata(requiredMetadata).
		WithMethods(
			map[string]server.MethodFunc{
				"hello":   ext.SayHello,
				"goodbye": ext.SayGoodbye,
			},
		).Build()
}
