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

func New()(*TokenServiceConn,error){

	grpcPortTarget:=fmt.Sprintf("%s:%s",os.Getenv("tokenservice"),os.Getenv("grpcPort"))
	
	conn,err:=grpc.Dial(grpcPortTarget,grpc.WithInsecure())
	
	if err != nil {
        log.Printf("Failed to connect to token service: %v", err)
		return nil,err
    }
    
	tokenServiceClientConn:=tokenRpcServerProtoFiles.NewTokenServiceClient(conn)

	tsc:=&TokenServiceConn{}

	tsc.TokenClient=tokenServiceClientConn

	return tsc,nil
}
