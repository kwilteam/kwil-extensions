package hello

import (
	"log"

	"github.com/kwilteam/kwil-extensions/server"
)

func NewHelloWorldExtension(logger *log.Logger) (*server.ExtensionServer, error) {
	ext := &HelloWorldExt{}

	return server.Builder().
		Named(ext.Name()).
		WithInitializer(initialize).
		WithMethods(
			map[string]server.MethodFunc{
				"hello":   ext.SayHello,
				"goodbye": ext.SayGoodbye,
			},
		).
		WithLoggerFunc(func(l string) {
			logger.Printf("log received from extension: %s", l)
		}).
		Build()
}
