package main

import (
	appRoutes "authenticationService/routes"
	"authenticationService/server"
	"authenticationService/store"
	tokenServiceConnPackage "authenticationService/tokenServiceConn"
	"errors"
	"os"
	"strings"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)


func main(){

	// *************http server implementation***************

	mongoUrl:=os.Getenv("mongoConnectionString")
	httpPort:=os.Getenv("httpPort")

	fmt.Println("connecting with token server over grpc!!")
	tsc,err:=tokenServiceConnPackage.New()

	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("connected with token server over grpc!!")

	fmt.Println("Creating store!!")
	
	store:=store.New(mongoUrl)

	fmt.Println("Creating server!!")
	server:=server.New(store)

	fmt.Println("Creating routes!!")
	routes:=appRoutes.InitRoutes(server,tsc)

	fmt.Println("starting server on "+httpPort)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Content-Disposition"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE"})
	
	err=http.ListenAndServe(fmt.Sprintf(":%s",httpPort),handlers.CORS(headersOk,originsOk,methodsOk)(routes))

	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("running server!!")
}

func getConnectionString(conenctString string) (string,error){
	startIndex := strings.Index(conenctString, "mongodb+srv:")
	if startIndex == -1 {
		return "",errors.New("MongoDB connection string not found")
	}

	// Trim the string to get the MongoDB connection string
	trimmed := conenctString[startIndex:]

	// Find the end position of the MongoDB connection string
	endIndex := strings.Index(trimmed, ",")
	if endIndex == -1 {
		fmt.Println("Invalid decoded string format")
		return "",errors.New("Invalid decoded string format")
	}

	// Extract the MongoDB connection string
	mongoUrl := trimmed[:endIndex]

	// Print the extracted MongoDB connection string
	return mongoUrl,nil
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

