package main

import (
	"authenticationService/authenticationServiceRpcServer"
	appRoutes "authenticationService/routes"
	"authenticationService/server"
	"authenticationService/store"
	"encoding/base64"

	// "context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/handlers"
	"google.golang.org/grpc"
)

func main(){

	secret:=os.Getenv("secretData")

	decodedBytes, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)
		return
	}

	// Convert decoded bytes to a string
	mongoUrl := string(decodedBytes)
	grpcPort:=os.Getenv("grpcPort")
	httpPort:=os.Getenv("httpPort")

	//**************grpc server implementation***************

	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		grpcServerInstance:=grpc.NewServer()
		authenticationServiceRpcServerType := &authenticationServiceRpcServer.ValidationServerGrpcStruct{}

		authenticationServiceRpcServer.RegisterTokenValidationServer(grpcServerInstance,authenticationServiceRpcServerType)

		listener, err := net.Listen("tcp", fmt.Sprintf(":%s",grpcPort))
		if err != nil {
			fmt.Printf("Failed to listen: %v", err)
			return
		}

		fmt.Printf("gRPC server listening on port %s\n", grpcPort)
		if err := grpcServerInstance.Serve(listener); err != nil {
			fmt.Printf("Failed to serve: %v", err)
		}
	}()

	//*************http server implementation***************
	
	go func(){
		defer wg.Done()

		fmt.Println("Creating store!!")
		
		store:=store.New(mongoUrl)
	
		fmt.Println("Creating server!!")
		server:=server.New(store)
	
		fmt.Println("Creating routes!!")
		routes:=appRoutes.InitRoutes(server)
	
		fmt.Println("starting server on "+httpPort)
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Content-Disposition"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE"})
		
		err:=http.ListenAndServe(fmt.Sprintf(":%s",httpPort),handlers.CORS(originsOk,headersOk,methodsOk)(routes))
	
		if(err!=nil){
			log.Fatal(err)
		}
	
		fmt.Println("running server!!")
	}()

	wg.Wait()

	
}

// func testGRPCConnection(grpcPort string) {
// 	// Set up a connection to the gRPC server.
// 	fmt.Println("i am called")


// 	conn, err := grpc.Dial(fmt.Sprintf(":%s",grpcPort), grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Failed to dial server: %v", err)
// 	}
// 	defer conn.Close()

// 	// Create a client for the authentication service.
// 	client := authenticationServiceRpcServer.NewTokenValidationClient(conn)

// 	// Call a method on the gRPC server.
// 	response, err := client.ValidateToken(context.Background(), &authenticationServiceRpcServer.TokenRequest{Token: "your token"})
// 	if err != nil {
// 		log.Fatalf("Error calling ValidateToken: %v", err)
// 	}

// 	// Print the response from the server.
// 	log.Printf("Response from server: %v\n", response)
// }

