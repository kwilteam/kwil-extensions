package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	gen "github.com/kwilteam/kwil-extensions/gen"
	"github.com/kwilteam/kwil-extensions/types"
	"github.com/kwilteam/kwil-extensions/types/convert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ExtensionClient struct {
	extClient gen.ExtensionServiceClient
	conn      *grpc.ClientConn

	// timeout is the timeout for all extension calls and the initial connection
	timeout time.Duration
}

func NewExtensionClient(ctx context.Context, target string, opts ...ClientOpt) (*ExtensionClient, error) {
	extClient := &ExtensionClient{
		timeout: time.Duration(5 * time.Second),
	}

	for _, opt := range opts {
		opt(extClient)
	}

	dialCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	dialOpts := extClient.grpcDialOpts()

	grpcClient, err := grpc.DialContext(dialCtx, target, dialOpts...)
	if err != nil {
		return nil, err
	}

	extClient.conn = grpcClient
	extClient.extClient = gen.NewExtensionServiceClient(grpcClient)

	return extClient, nil
}

func (c *ExtensionClient) Close() error {
	return c.conn.Close()
}

func (c *ExtensionClient) GetName(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := c.extClient.Name(ctx, &gen.NameRequest{})
	if err != nil {
		return "", err
	}

	return resp.Name, nil
}

func (c *ExtensionClient) ListMethods(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := c.extClient.ListMethods(ctx, &gen.ListMethodsRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Methods, nil
}

func (c *ExtensionClient) CallMethod(execCtx *types.ExecutionContext, method string, args ...any) ([]any, error) {
	ctx, cancel := context.WithTimeout(execCtx.Ctx, c.timeout)
	defer cancel()

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

	resp, err := c.extClient.Execute(ctx, &gen.ExecuteRequest{
		Name:     strings.ToLower(method),
		Args:     pbArgs,
		Metadata: execCtx.Metadata,
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
		outputs = append(outputs, scalarOutput.Value)
	}

	return outputs, nil
}

func (c *ExtensionClient) Initialize(ctx context.Context, metadata map[string]string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := c.extClient.Initialize(ctx, &gen.InitializeRequest{
		Metadata: metadata,
	})
	if err != nil {
		return nil, err
	}

	return resp.Metadata, nil
}

func (c *ExtensionClient) grpcDialOpts() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
}
