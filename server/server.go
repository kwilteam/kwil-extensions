package server

import (
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	gen "github.com/kwilteam/kwil-extensions/gen"
	"google.golang.org/grpc"
)

// ExtensionServer is a server that runs an extension
type ExtensionServer struct {
	extension *Extension

	logFn      LoggerFunc
	grpcServer *grpc.Server
}

// LoggerFunc is a function that consumes logs from the Extension Server.
type LoggerFunc func(l string)

func (e *ExtensionServer) Serve(lis net.Listener) error {
	if e.started() {
		return fmt.Errorf("server already started")
	}

	e.grpcServer = e.newGrpcServer()

	gen.RegisterExtensionServiceServer(e.grpcServer, e.extension)

	return e.grpcServer.Serve(lis)
}

func (e *ExtensionServer) GracefulStop() error {
	if !e.started() {
		return fmt.Errorf("server has not beed started")
	}

	e.grpcServer.GracefulStop()
	e.grpcServer = nil

	return nil
}

func (e *ExtensionServer) started() bool {
	return e.grpcServer != nil
}

func (e *ExtensionServer) newGrpcServer() *grpc.Server {
	recoveryFunc := func(p interface{}) error {
		e.logFn(fmt.Sprintf("panic: %v", p))
		return nil
	}
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(recoveryFunc),
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
		),
	)

	return server
}
