package v1

import (
	healthcheck "authenticationService/server/controller/v1/health_check"

	"github.com/gorilla/mux"
)

func V1() *mux.Router{
	router:=mux.NewRouter()
	
	router.HandleFunc("/health",healthcheck.HealthCheckController).Methods("GET")
	
	return router
}