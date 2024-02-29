package v1

import (
	"authenticationService/server"
	"authenticationService/server/v1/controller/health_check"
	signUp "authenticationService/server/v1/controller/sign_up"

	"github.com/gorilla/mux"
)


func V1(server *server.Server) *mux.Router{
	router:=mux.NewRouter()
	
	router.HandleFunc("/health",health_check.GetHealth(server)).Methods("GET")
	router.HandleFunc("/sign_up",signUp.SignUp(server)).Methods("GET")

	return router
}