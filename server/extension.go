package server

import (
	"github.com/kwilteam/kwil-extensions/types"
)

type MethodFunc func(ctx *types.ExecutionContext, inputs ...*types.ScalarValue) ([]*types.ScalarValue, error)

type ExtensionConfig struct {
	Name             string
	RequiredMetadata map[string]string
	Methods          map[string]MethodFunc
}
