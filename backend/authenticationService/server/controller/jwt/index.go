package jwt

import (
	"authenticationService/server"
	"authenticationService/server/controller/jwt/jwt_v1"
	"authenticationService/utils"

	"github.com/gorilla/mux"
)

func ValidateJwtInitRoute(server *server.Server)*mux.Router{
	router:=mux.NewRouter()

	utils.RouterUtils(router,"/v1",jwt_v1.ValidateJwtV1Routes(server))

	return router
}