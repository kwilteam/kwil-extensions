package extension

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/kwilteam/kwil-extensions/client"
	"github.com/kwilteam/kwil-extensions/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFractalExt_grpc(t *testing.T) {
	clt, err := client.NewExtensionClient(context.Background(), "127.0.0.1:50051", client.WithTimeout(100*time.Second))
	assert.NoError(t, err)

	methods, err := clt.ListMethods(context.Background())
	assert.NoError(t, err)
	fmt.Println(methods)

	metadata := map[string]string{
		"registry_address": "0x274b028b03A250cA03644E6c578D81f019eE1323",
		"chain_name":       "goerli",
	}
	ctx := &types.ExecutionContext{
		Ctx:      context.Background(),
		Metadata: metadata,
	}

	// Output:
	res, err := clt.CallMethod(ctx, "is_user_in_list",
		"e55149bfd05867a51672a24235e3511767bd64cb1b250c33da303d5be58d2bdd", "plus")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.EqualValues(t, 1, res[0])

	res, err = clt.CallMethod(ctx, "grants_for")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.EqualValues(t, 3, res[0])
}

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

			//err = e.Configure(map[string]string{
			//	"registry_address": tt.fields.RegistryAddress,
			//})
			//assert.NoError(ttt, err)

			fmt.Println("registry address:", tt.fields.RegistryAddress)
			fmt.Println("wallet address:", testAddress)
			fmt.Println("list name:", testListName)
			ctx := &types.ExecutionContext{
				Ctx: context.Background(),
				Metadata: map[string]string{
					"registry_address": tt.fields.RegistryAddress,
					"chain_name":       "goerli",
				},
			}
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

			//// only for ganache
			//got, err = srv.Methods["grants_for"](ctx)
			//assert.NoError(ttt, err)
			//fmt.Println("grants for:", got[0].Int())
		})
	}
}
