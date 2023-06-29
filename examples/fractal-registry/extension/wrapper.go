package extension

import (
	"fmt"
	"github.com/kwilteam/kwil-extensions/server"
	"github.com/kwilteam/kwil-extensions/types"
)

func WithInputOutputCheck(
	fn server.MethodFunc,
	argsTypeList []types.ScalarType,
	returnTypeList []types.ScalarType) server.MethodFunc {
	return func(ctx *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
		if len(values) != len(argsTypeList) {
			return nil, ErrInvalidArgumentNum{
				Expect: len(argsTypeList),
				Got:    len(values),
			}
		}

		for i, v := range values {
			if v.Type != argsTypeList[i] {
				return nil, ErrInvalidArgumentType{
					Expect: argsTypeList[i],
					Got:    v.Type,
					Pos:    i,
				}
			}
		}

		res, err := fn(ctx, values...)
		if err != nil {
			return nil, fmt.Errorf("execution got err: %w", err)
		}

		if len(res) != len(returnTypeList) {
			return nil, ErrInvalidReturnNum{
				Expect: len(returnTypeList),
				Got:    len(res),
			}
		}

		for i, v := range res {
			if v.Type != returnTypeList[i] {
				return nil, ErrInvalidReturnType{
					Expect: returnTypeList[i],
					Got:    v.Type,
					Pos:    i,
				}
			}
		}

		return res, nil
	}
}
