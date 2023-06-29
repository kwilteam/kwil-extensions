package builder

import (
	"github.com/kwilteam/kwil-extensions/server"
)

func Builder() ExtensionNameBuilder {
	return &ExtensionBuilder{
		config: &server.ExtensionConfig{},
	}
}

type ExtensionBuilder struct {
	config *server.ExtensionConfig
}

type ExtensionNameBuilder interface {
	Named(string) ExtensionMetadataBuilder
}

type ExtensionMetadataBuilder interface {
	WithRequiredMetadata(map[string]string) ExtensionMethodBuilder
}

type ExtensionMethodBuilder interface {
	WithMethods(map[string]server.MethodFunc) ExtensionBuildBuilder
}

type ExtensionBuildBuilder interface {
	Build() (*server.Server, error)
}

func (b *ExtensionBuilder) Named(name string) ExtensionMetadataBuilder {
	b.config.Name = name
	return b
}

func (b *ExtensionBuilder) WithRequiredMetadata(requiredMetadata map[string]string) ExtensionMethodBuilder {
	b.config.RequiredMetadata = requiredMetadata
	return b
}

func (b *ExtensionBuilder) WithMethods(methods map[string]server.MethodFunc) ExtensionBuildBuilder {
	b.config.Methods = methods
	return b
}

func (b *ExtensionBuilder) Build() (*server.Server, error) {
	return server.NewExtensionServer(b.config)
}
