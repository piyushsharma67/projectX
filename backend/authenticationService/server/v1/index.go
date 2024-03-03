package v1

import (
	"authenticationService/server"
	"authenticationService/server/v1/controller/health_check"
	"authenticationService/server/v1/controller/login"
	signUp "authenticationService/server/v1/controller/sign_up"

	"github.com/gorilla/mux"
)


func V1(server *server.Server) *mux.Router{
	router:=mux.NewRouter()
	
	router.HandleFunc("/health",health_check.GetHealth(server)).Methods("GET")
	router.HandleFunc("/sign_up",signUp.SignUpController(server)).Methods("POST")
	router.HandleFunc("/fetch_user_by_email",signUp.FetchUserByEmailController(server)).Methods("GET")
	router.HandleFunc("/drop_user_table",signUp.DropUserTable(server)).Methods("POST")
	router.HandleFunc("/login",login.Login(server)).Methods("GET")

	return router
}