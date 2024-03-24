package main

import (
	"net"
	"os"
	"tokenService/tokenRpcServerProtoFiles"

	"fmt"

	"google.golang.org/grpc"
)

func main(){

	grpcPort:=os.Getenv("grpcPort")
	
	// **************grpc server implementation***************

	grpcServerInstance:=grpc.NewServer()
	tokenServiceGrpcType := &tokenRpcServerProtoFiles.TokenServices{}

	tokenRpcServerProtoFiles.RegisterTokenValidationServer(grpcServerInstance,tokenServiceGrpcType)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s",grpcPort))
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	fmt.Printf("gRPC server listening on port %s\n", grpcPort)
	if err := grpcServerInstance.Serve(listener); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
	
}
