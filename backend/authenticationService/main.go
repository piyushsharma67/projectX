package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	port := os.Getenv("PORT")

	// user:="root"
	// password:=os.Getenv("MYSQL_ROOT_PASSWORD")
	// database:=os.Getenv("MYSQL_DATABASE")
	// databasePort:=os.Getenv("MY_SQL_PORT")
	fmt.Println("Creating store!!")
	
	// store:=store.New(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,password,"sql-service",databasePort,database))

	fmt.Println("Creating server!!")
	// server:=server.New(store)

	fmt.Println("Creating routes!!")
	// routes:=appRoutes.InitRoutes(server)

	fmt.Println("starting server on "+port)
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Content-Disposition"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE"})
	http.HandleFunc("/health",ServerHealth)
	err:=http.ListenAndServe(fmt.Sprintf(":%s","3004"),nil)

	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("running server!!")

}

func ServerHealth(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("My health is Ok!!"))
}