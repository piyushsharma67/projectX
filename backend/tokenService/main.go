package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"tokenService/tokenRpcServerProtoFiles"

	"fmt"

	"google.golang.org/grpc"
)

func main(){

	grpcPort:=os.Getenv("grpcPort")
	
	// **************grpc server implementation***************
	
	grpcServerInstance:=grpc.NewServer()
	tokenServiceGrpcType := &tokenRpcServerProtoFiles.TokenServices{}

	tokenRpcServerProtoFiles.RegisterTokenServiceServer(grpcServerInstance,tokenServiceGrpcType)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s",grpcPort))
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	go func() {
		fmt.Printf("gRPC server listening on port %s\n", grpcPort)
		if err := grpcServerInstance.Serve(listener); err != nil {
			fmt.Printf("Failed to serve: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Block until a signal is received

	fmt.Println("Shutting down server...")
	grpcServerInstance.GracefulStop()
	fmt.Println("Server stopped gracefully")
}
