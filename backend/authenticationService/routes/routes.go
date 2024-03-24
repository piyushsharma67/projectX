package appRoutes

import (
	"authenticationService/server"
	"authenticationService/server/controller"

	"authenticationService/server/controller/authenticate"
	utils "authenticationService/utils"

	"github.com/gorilla/mux"
)

func InitRoutes(server *server.Server) *mux.Router{
	router:=mux.NewRouter()

	utils.RouterUtils(router,"/auth/authenticate",authenticate.AuthenticateInitRoute(server))
	router.HandleFunc("/auth/health",controller.ServerHealth).Methods("GET")
	
	return router
}