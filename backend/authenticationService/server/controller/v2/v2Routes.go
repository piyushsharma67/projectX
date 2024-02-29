package v2

import (
	"authenticationService/server"
	healthcheck "authenticationService/server/controller/v2/health_check"

	"github.com/gorilla/mux"
)

func V2(server *server.Server) *mux.Router{
	router:=mux.NewRouter()
	
	router.HandleFunc("/health",healthcheck.HealthCheckController(server)).Methods("GET")
	
	return router
}