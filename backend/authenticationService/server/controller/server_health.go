package controller

import "net/http"


func ServerHealth(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("Hi am in Goodhealth"))
}