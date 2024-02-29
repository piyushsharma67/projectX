package health_check

import (
	"authenticationService/server"
	"net/http"
)

func  HealthCheckController(server *server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("hi i am working V1 health"))
	}
}