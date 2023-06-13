package builder

import (
	"github.com/kwilteam/kwil-extensions/server"
	"github.com/kwilteam/kwil-extensions/types"
)

func Builder() *ExtensionBuilder {
	return &ExtensionBuilder{}
}

type ExtensionBuilder struct {
	config *server.ExtConf
}

type ExtensionConfigFuncBuilder interface {
	WithConfigFunc(func(map[string]string) error) ExtensionMetadataBuilder
}

type ExtensionMetadataBuilder interface {
	WithRequiredMetadata(map[string]string) ExtensionMethodBuilder
}

type ExtensionMethodBuilder interface {
	WithMethods(...func([]*types.ScalarValue, map[string]string) ([]*types.ScalarValue, error)) ExtensionBuilder
}
