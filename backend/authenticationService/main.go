package main

import (
	appRoutes "authenticationService/routes"
	"authenticationService/server"
	"authenticationService/store"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main(){
	port := os.Getenv("PORT")

	user:="root"
	password:=os.Getenv("MYSQL_ROOT_PASSWORD")
	database:=os.Getenv("MYSQL_DATABASE")
	databasePort:=os.Getenv("MY_SQL_PORT")
	fmt.Println("Creating store!!")
	
	store:=store.New(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,password,"sql-service",databasePort,database))

	fmt.Println("Creating server!!")
	server:=server.New(store)

	fmt.Println("Creating routes!!")
	routes:=appRoutes.InitRoutes(server)

	fmt.Println("starting server on "+port)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Content-Disposition"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE"})

	err:=http.ListenAndServe(fmt.Sprintf(":%s",port),handlers.CORS(originsOk, headersOk, methodsOk)(routes))

	if(err!=nil){
		log.Fatal("error occured")
	}

}