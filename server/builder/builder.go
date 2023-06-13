package builder

import (
	"github.com/kwilteam/kwil-extensions/server"
)

func Builder() ExtensionConfigFuncBuilder {
	return &ExtensionBuilder{
		config: &server.ExtensionConfig{},
	}
}

type ExtensionBuilder struct {
	config *server.ExtensionConfig
}

type ExtensionConfigFuncBuilder interface {
	WithConfigFunc(server.ConfigFunc) ExtensionMetadataBuilder
}

type ExtensionMetadataBuilder interface {
	WithRequiredMetadata(map[string]string) ExtensionMethodBuilder
}

type ExtensionMethodBuilder interface {
	WithMethods(...server.MethodFunc) ExtensionBuildBuilder
}

type ExtensionBuildBuilder interface {
	Build() (*server.Server, error)
}

func (b *ExtensionBuilder) WithConfigFunc(configFunc server.ConfigFunc) ExtensionMetadataBuilder {
	b.config.ConfigFunc = configFunc
	return b
}

func (b *ExtensionBuilder) WithRequiredMetadata(requiredMetadata map[string]string) ExtensionMethodBuilder {
	b.config.RequiredMetadata = requiredMetadata
	return b
}

func (b *ExtensionBuilder) WithMethods(methods ...server.MethodFunc) ExtensionBuildBuilder {
	b.config.Methods = methods
	return b
}

func (b *ExtensionBuilder) Build() (*server.Server, error) {
	return server.NewExtensionServer(b.config)
}
