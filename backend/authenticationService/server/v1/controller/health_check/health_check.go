package health_check

import (
	"authenticationService/server"
	"net/http"
)

func GetHealth(server *server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Hi i am V1 health checker"))
	}
}