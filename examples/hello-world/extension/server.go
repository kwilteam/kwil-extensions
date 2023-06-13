package hello

import (
	"github.com/kwilteam/kwil-extensions/server"
	"github.com/kwilteam/kwil-extensions/server/builder"
)

func NewHelloWorldExtension() (*server.Server, error) {
	ext := &HelloWorldExt{}

	return builder.Builder().
		WithConfigFunc(ext.Configure).
		WithRequiredMetadata(requiredMetadata).
		WithMethods(
			ext.SayHello,
			ext.SayGoodbye,
		).Build()
}
