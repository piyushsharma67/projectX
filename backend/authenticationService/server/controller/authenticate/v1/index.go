package v1

import (
	"authenticationService/server"
	tokenServiceConnPackage "authenticationService/tokenServiceConn"

	"github.com/gorilla/mux"
)

func V1(server *server.Server,tsc *tokenServiceConnPackage.TokenServiceConn) *mux.Router{
	router:=mux.NewRouter()
	
	router.HandleFunc("/signup",SignUpController(server,tsc)).Methods("POST")
	router.HandleFunc("/login",LoginController(server,tsc)).Methods("POST")

	return router
}