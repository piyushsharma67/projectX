package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Printf("starting server on port 8080")

	err:=http.ListenAndServe(":8080",nil)

	if(err!=nil){
		log.Fatal("error occured")
	}
}