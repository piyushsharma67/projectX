package v1

import (
	"authenticationService/server"

	"github.com/gorilla/mux"
)


func V1(server *server.Server) *mux.Router{
	router:=mux.NewRouter()
	
	router.HandleFunc("/signup",SignUpController(server)).Methods("POST")

	return router
}