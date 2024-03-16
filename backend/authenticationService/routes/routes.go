package appRoutes

import (
	"authenticationService/server"
	"authenticationService/server/controller/authenticate"
	"authenticationService/server/controller/jwt"
	utils "authenticationService/utils"

	"github.com/gorilla/mux"
)

func InitRoutes(server *server.Server) *mux.Router{
	router:=mux.NewRouter()

	utils.RouterUtils(router,"/authenticate",authenticate.AuthenticateInitRoute(server))
	utils.RouterUtils(router,"/jwt",jwt.ValidateJwtInitRoute(server))
	
	return router
}