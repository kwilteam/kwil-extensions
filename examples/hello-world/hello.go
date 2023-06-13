package helloworld

import (
	"fmt"

	"github.com/kwilteam/kwil-extensions/server"
)

type HelloWorldExt struct {
	server *server.Server

	greeting    string
	punctuation string
}

func NewHelloWorldExtension() *HelloWorldExt {
	ext := &HelloWorldExt{}

	server := server.NewExtensionServer(&helloWorldExtension, ext)

}

func (e *HelloWorldExt) Configure(newConfig map[string]string) error {
	// TODO: change this
	greeting, ok := config["greeting"]
	if !ok {
		return fmt.Errorf("greeting is required")
	}

	punctuation, ok := config["punctuation"]
	if !ok {
		return fmt.Errorf("punctuation is required")
	}

	e.greeting = greeting
	e.punctuation = punctuation

	return nil
}

func (h *HelloWorldExt) sayHello(name string) string {

	return fmt.Sprintf("%s %s%s", h.greeting, name, h.punctuation)
}
