package main

import (
	appRoutes "authenticationService/routes"
	"authenticationService/server"
	"authenticationService/store"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	port := os.Getenv("PORT")

	user:=os.Getenv("MYSQL_USER")
	password:=os.Getenv("MYSQL_PASSWORD")
	database:=os.Getenv("MYSQL_DATABASE")

	fmt.Println("Creating store!!")
	store:=store.New(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,password,"localhost",port,database))

	fmt.Println("Creating server!!")
	server:=server.New(store)

	fmt.Println("Creating routes!!")
	routes:=appRoutes.InitRoutes(server)

	fmt.Println("starting server on "+port)

	err:=http.ListenAndServe(":"+port,routes)

	if(err!=nil){
		log.Fatal("error occured")
	}
}