package client

import (
	"context"
	"fmt"
	"strings"

	gen "github.com/kwilteam/kwil-extensions/gen"
	"github.com/kwilteam/kwil-extensions/types"
	"github.com/kwilteam/kwil-extensions/types/convert"
	"google.golang.org/grpc"
)

type ExtensionClient struct {
	extClient gen.ExtensionServiceClient
	conn      *grpc.ClientConn
}

func NewExtensionClient(target string, opts ...grpc.DialOption) (*ExtensionClient, error) {
	grpcClient, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, err
	}

	return &ExtensionClient{
		extClient: gen.NewExtensionServiceClient(grpcClient),
		conn:      grpcClient,
	}, nil
}

func (c *ExtensionClient) Close() error {
	return c.conn.Close()
}

func (c *ExtensionClient) Configure(ctx context.Context, config map[string]string) error {
	success, err := c.extClient.Configure(ctx, &gen.ConfigureRequest{
		Config: config,
	})
	if err != nil {
		return err
	}

	if !success.Success {
		return fmt.Errorf("failed to configure extension. the extension did not provide a reason")
	}

	return nil
}

func (c *ExtensionClient) ListMethods(ctx context.Context) ([]string, error) {
	resp, err := c.extClient.ListMethods(ctx, &gen.ListMethodsRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Methods, nil
}

func (c *ExtensionClient) CallMethod(ctx *types.ExecutionContext, method string, args ...any) ([]any, error) {
	var encodedArgs []*types.ScalarValue
	for _, arg := range args {
		scalarVal, err := types.NewScalarValue(arg)
		if err != nil {
			return nil, fmt.Errorf("error encoding argument: %s", err.Error())
		}

		encodedArgs = append(encodedArgs, scalarVal)
	}

	pbArgs, err := convert.ConvertScalarToPb(encodedArgs)
	if err != nil {
		return nil, fmt.Errorf("error converting arguments: %s", err.Error())
	}

	resp, err := c.extClient.Execute(ctx.Ctx, &gen.ExecuteRequest{
		Name:     strings.ToLower(method),
		Args:     pbArgs,
		Metadata: ctx.Metadata,
	})
	if err != nil {
		return nil, err
	}

	scalarOutputs, err := convert.ConvertScalarFromPb(resp.Outputs)
	if err != nil {
		return nil, fmt.Errorf("error converting outputs: %s", err.Error())
	}

	var outputs []any
	for _, scalarOutput := range scalarOutputs {
		outputs = append(outputs, scalarOutput.Any())
	}

	return outputs, nil
}

func (c *ExtensionClient) GetMetadata(ctx context.Context) (map[string]string, error) {
	resp, err := c.extClient.GetMetadata(ctx, &gen.GetMetadataRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Metadata, nil
}
