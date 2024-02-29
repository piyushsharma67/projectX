package v2

import (
	"authenticationService/server"
	"authenticationService/server/v2/controller/health_check"

	"github.com/gorilla/mux"
)

func V2(server *server.Server) *mux.Router{
	router:=mux.NewRouter()
	
	router.HandleFunc("/health",health_check.GetHealth(server)).Methods("GET")
	
	return router
}