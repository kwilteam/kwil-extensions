package main

import (
	"github.com/caarlos0/env/v8"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/kwilteam/extension-fractal-demo/extension"
	gen "github.com/kwilteam/kwil-extensions/gen"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

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

func main() {
	cfg := extension.Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	ext, err := extension.NewFractalExt(cfg.RpcUrl)
	if err != nil {
		panic(err)
	}

	srv, err := ext.BuildServer()
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "fractal:", log.LstdFlags)

	lis, err := net.Listen("tcp", cfg.ListenAddr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSvr := newGrpcServer(logger)
	gen.RegisterExtensionServiceServer(grpcSvr, srv)

	go func() {
		logger.Printf("listening on %s", cfg.ListenAddr())
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
