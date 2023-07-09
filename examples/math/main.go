package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kwilteam/math/extension"
)

const (
	listenAddress = ":50051"
)

func main() {
	logger := log.New(os.Stdout, "math: ", log.LstdFlags)

	svr, err := extension.NewMathExtension(logger)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		logger.Printf("listening on %s", listenAddress)
		if err := svr.Serve(lis); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Printf("\nshutting down")

	err = svr.GracefulStop()
	if err != nil {
		logger.Fatalf("failed to shutdown extension server: %v", err)
	}

	logger.Printf("shutdown complete")
}
