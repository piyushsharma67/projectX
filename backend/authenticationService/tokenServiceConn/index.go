package tokenServiceConnPackage

import (
	"authenticationService/authRpcServerProtoFiles"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

type TokenServiceConn struct{
	TokenClient authRpcServerProtoFiles.TokenServiceClient
}

const (
    tokenServiceAddr = "tokenservice:5002" // Assuming token service listens on port 5001
)

func New()(*TokenServiceConn,error){
	
	conn,err:=grpc.Dial(tokenServiceAddr,grpc.WithInsecure())
	
	if err != nil {
        log.Printf("Failed to connect to token service: %v", err)
		return nil,err
    }
    
	tokenServiceClientConn:=authRpcServerProtoFiles.NewTokenServiceClient(conn)

	tsc:=&TokenServiceConn{}

	tsc.TokenClient=tokenServiceClientConn

	log.Printf("Connected to token service at %s", tokenServiceAddr)

	err=pingTokenService("5002")

	if(err!=nil){
		fmt.Println("nmn"+err.Error())
	}

	return tsc,nil
}

func pingTokenService(grpcPort string) error {
	// Set up a connection to the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // Timeout after 1 second
	defer cancel()

	
	conn, err := grpc.DialContext(ctx, fmt.Sprintf(":%s", "5002"), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial tokenservice: %v", err)
	}
	defer conn.Close()

	// Create a client for the tokenservice.
	client := authRpcServerProtoFiles.NewTokenServiceClient(conn)

	// Call a method on the gRPC server.
	_, err = client.GenerateToken(context.Background(), &authRpcServerProtoFiles.GenerateTokenRequest{})
    if err != nil {
        return fmt.Errorf("Error calling GenerateToken: %v", err)
    }

	return err
}