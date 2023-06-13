package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	hello "github.com/kwilteam/extension-hello-world/extension"
	gen "github.com/kwilteam/kwil-extensions/gen"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
)

const (
	listenAddress = ":50051"
)

func main() {
	svr, err := hello.NewHelloWorldExtension()
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "hello-world: ", log.LstdFlags)

	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSvr := newGrpcServer(logger)
	gen.RegisterExtensionServiceServer(grpcSvr, svr)

	go func() {
		logger.Printf("listening on %s", listenAddress)
		if err := grpcSvr.Serve(lis); err != nil {
			logger.Printf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Printf("\nshutting down")

	grpcSvr.GracefulStop()

	logger.Printf("shutdown complete")
}

func newGrpcServer(logger *log.Logger) *grpc.Server {

	recoveryFunc := func(p interface{}) error {
		logger.Printf("panic: %v", p)
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
