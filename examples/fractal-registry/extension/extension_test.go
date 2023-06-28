package extension

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/kwilteam/kwil-extensions/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFractalExt_Behavior(t *testing.T) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		t.Errorf("parse config failed: %v", err)
	}
	if cfg.RpcUrl == "" {
		t.Errorf("rpc url is empty")
	}

	type fields struct {
		RegistryAddress string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "goerli",
			fields: fields{
				RegistryAddress: "0x4D9DE1bb481B9dA37A7a7E3a07F6f60654fEe7BB",
			},
		},
		//{
		//	name: "ganache",
		//	fields: fields{
		//		RegistryAddress: "0x274b028b03A250cA03644E6c578D81f019eE1323",
		//	},
		//},
	}

	const testAddress = "0x640568976c2CDc8789E44B39369D5Bc44B1e6Ad7"
	const testListName = "plus"

	for _, tt := range tests {
		t.Run(tt.name, func(ttt *testing.T) {
			e, err := NewFractalExt(cfg.RpcUrl)
			assert.NoError(ttt, err)

			srv, err := e.BuildServer()
			assert.NoError(ttt, err)

			err = e.Configure(map[string]string{
				"registry_address": tt.fields.RegistryAddress,
			})
			assert.NoError(ttt, err)

			fmt.Println("registry address:", e.RegistryAddress.String())
			fmt.Println("wallet address:", testAddress)
			fmt.Println("list name:", testListName)
			ctx := &types.ExecutionContext{Ctx: context.Background()}
			//
			getFractalIDArgs, err := encodeScalarValues(testAddress)
			assert.NoError(ttt, err)
			got, err := srv.Methods["get_fractal_id"](ctx, getFractalIDArgs...)
			assert.NoError(ttt, err)
			fmt.Println("got fractal id:", got[0].String())

			//
			fractalID := got[0].String()
			isUserInListArgs, err := encodeScalarValues(fractalID, testListName)
			assert.NoError(ttt, err)
			got, err = srv.Methods["is_user_in_list"](ctx, isUserInListArgs...)
			assert.NoError(ttt, err)
			fmt.Println("is user in list:", got[0].Int())

			//
			got, err = srv.Methods["grants_for"](ctx)
			assert.NoError(ttt, err)
			fmt.Println("grants for:", got[0].Int())
		})
	}
}
