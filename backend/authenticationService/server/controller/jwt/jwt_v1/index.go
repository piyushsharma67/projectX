package jwt_v1

import (
	"authenticationService/server"

	"github.com/gorilla/mux"
)

func ValidateJwtV1Routes(server *server.Server) *mux.Router{
	router:=mux.NewRouter()

	router.HandleFunc("/validate",ValidateJwtController(server))
	return router
}