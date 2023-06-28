package extension

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kwilteam/extension-fractal-demo/extension/registry"
	"github.com/kwilteam/kwil-extensions/server"
	"github.com/kwilteam/kwil-extensions/server/builder"
	"github.com/kwilteam/kwil-extensions/types"
)

type FractalExt struct {
	RPCURL          string
	eth             *ethclient.Client
	RegistryAddress common.Address
	Registry        *registry.Registry
}

func NewFractalExt(rpcURL string) (*FractalExt, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("dial rpc failed: %w", err)
	}

	ext := &FractalExt{
		RPCURL: rpcURL,
		eth:    client,
	}

	blockNum, err := ext.GetBlockHeight(
		&types.ExecutionContext{
			Ctx: context.Background(),
		},
		[]*types.ScalarValue{}...)
	if err != nil {
		return nil, fmt.Errorf("get block height failed: %w", err)
	}
	fmt.Printf("block height: %d\n", blockNum[0].Int())

	return ext, nil
}

func (e *FractalExt) BuildServer() (*server.Server, error) {
	return builder.Builder().
		WithConfigFunc(e.Configure).
		WithRequiredMetadata(requiredMetadata).
		WithMethods(
			map[string]server.MethodFunc{
				"get_block_height": WithInputOutputCheck(
					e.GetBlockHeight,
					[]types.ScalarType{},
					[]types.ScalarType{types.ScalarType_INT}),
				"get_fractal_id": WithInputOutputCheck(
					e.GetFractalID,
					[]types.ScalarType{types.ScalarType_STRING},
					[]types.ScalarType{types.ScalarType_STRING}),
				"is_user_in_list": WithInputOutputCheck(
					e.IsUserInList,
					[]types.ScalarType{types.ScalarType_STRING, types.ScalarType_STRING},
					[]types.ScalarType{types.ScalarType_INT}),
				"grants_for": WithInputOutputCheck(
					e.GrantsFor,
					[]types.ScalarType{},
					[]types.ScalarType{types.ScalarType_INT}),
			}).Build()
}

func (e *FractalExt) Configure(conf map[string]string) error {
	registryAddressStr, ok := conf["registry_address"]
	if !ok {
		return fmt.Errorf("registry_address is not set")
	}

	e.RegistryAddress = common.HexToAddress(registryAddressStr)

	instance, err := registry.NewRegistry(e.RegistryAddress, e.eth)
	if err != nil {
		return fmt.Errorf("create registry failed: %w", err)
	}
	e.Registry = instance

	return nil
}

func (e *FractalExt) GetBlockHeight(ctx *types.ExecutionContext, _ ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	num, err := e.eth.BlockNumber(ctx.Ctx)
	if err != nil {
		return nil, fmt.Errorf("get block number failed: %w", err)
	}

	return encodeScalarValues(num)
}

func (e *FractalExt) GetFractalID(_ *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	walletAddr := values[0].String()
	fractalID, err := e.Registry.GetFractalId(&bind.CallOpts{}, common.HexToAddress(walletAddr))
	if err != nil {
		return nil, fmt.Errorf("get fractal id failed: %w", err)
	}

	fractalIDStr := hex.EncodeToString(fractalID[:])
	return encodeScalarValues(fractalIDStr)
}

func (e *FractalExt) IsUserInList(_ *types.ExecutionContext, values ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	fractalID := values[0].String()
	listID := values[1].String()

	fractalIDByte, err := hex.DecodeString(fractalID)
	if err != nil {
		return nil, fmt.Errorf("decode fractal id failed: %w", err)
	}

	fractalIDByte32 := *abi.ConvertType(fractalIDByte, new([32]byte)).(*[32]byte)

	presence, err := e.Registry.IsUserInList(&bind.CallOpts{}, fractalIDByte32, listID)
	if err != nil {
		return nil, fmt.Errorf("get fractal id failed: %w", err)
	}

	// use int8 to represent bool
	var exist int8
	if presence {
		exist = 1
	} else {
		exist = 0
	}

	return encodeScalarValues(exist)
}

func (e *FractalExt) GrantsFor(_ *types.ExecutionContext, _ ...*types.ScalarValue) ([]*types.ScalarValue, error) {
	grantList, err := e.Registry.GrantsFor(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("get grants for failed: %w", err)
	}

	//var exist int8
	//if len(grantList) > 0 {
	//	exist = 1
	//} else {
	//	exist = 0
	//}

	return encodeScalarValues(len(grantList))
}

func encodeScalarValues(values ...any) ([]*types.ScalarValue, error) {
	scalarValues := make([]*types.ScalarValue, len(values))
	for i, v := range values {
		scalarValue, err := types.NewScalarValue(v)
		if err != nil {
			return nil, fmt.Errorf("convert value to scalar failed: %w", err)
		}

		scalarValues[i] = scalarValue
	}

	return scalarValues, nil
}

var requiredMetadata = map[string]string{}
