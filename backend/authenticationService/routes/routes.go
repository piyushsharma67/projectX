package appRoutes

import (
	"authenticationService/server"
	"authenticationService/server/controller"
	tokenServiceConnPackage "authenticationService/tokenServiceConn"

	"authenticationService/server/controller/authenticate"
	utils "authenticationService/utils"

	"github.com/gorilla/mux"
)

func InitRoutes(server *server.Server,tsc *tokenServiceConnPackage.TokenServiceConn) *mux.Router{
	router:=mux.NewRouter()

	utils.RouterUtils(router,"/auth/authenticate",authenticate.AuthenticateInitRoute(server,tsc))
	router.HandleFunc("/auth/health",controller.ServerHealth).Methods("GET")
	
	return router
}