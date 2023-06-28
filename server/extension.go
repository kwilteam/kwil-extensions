package server

import (
	"github.com/kwilteam/kwil-extensions/types"
)

type ConfigFunc func(map[string]string) error

type MethodFunc func(ctx *types.ExecutionContext, inputs ...*types.ScalarValue) ([]*types.ScalarValue, error)

type ExtensionConfig struct {
	ConfigFunc       ConfigFunc
	RequiredMetadata map[string]string
	Methods          map[string]MethodFunc
}
