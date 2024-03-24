package tokenServiceConnPackage

import (
	"authenticationService/tokenRpcServerProtoFiles"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
)

type TokenServiceConn struct{
	TokenClient tokenRpcServerProtoFiles.TokenServiceClient
}

const (
    tokenServiceAddr = "tokenservice" // Assuming token service listens on port 5001
)

func New()(*TokenServiceConn,error){

	grpcPortTarget:=fmt.Sprintf("tokenservice:%s",os.Getenv("grpcPort"))
	
	conn,err:=grpc.Dial(grpcPortTarget,grpc.WithInsecure())
	
	if err != nil {
        log.Printf("Failed to connect to token service: %v", err)
		return nil,err
    }
    
	tokenServiceClientConn:=tokenRpcServerProtoFiles.NewTokenServiceClient(conn)

	tsc:=&TokenServiceConn{}

	tsc.TokenClient=tokenServiceClientConn

	log.Printf("Connected to token service at %s", tokenServiceAddr)

	if(err!=nil){
		fmt.Println("nmnn"+err.Error())
	}

	return tsc,nil
}
