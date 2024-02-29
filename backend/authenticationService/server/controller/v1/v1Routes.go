package v1

import (
	"authenticationService/server"
	"authenticationService/server/controller/v1/health_check"

	"github.com/gorilla/mux"
)


func V1(server *server.Server) *mux.Router{
	router:=mux.NewRouter()
	
	router.HandleFunc("/health",health_check.HealthCheckController(server)).Methods("GET")
	
	return router
}