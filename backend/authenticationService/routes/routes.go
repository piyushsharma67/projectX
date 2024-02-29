package appRoutes

import (
	"authenticationService/server"
	v1 "authenticationService/server/v1"
	v2 "authenticationService/server/v2"

	utils "authenticationService/utils"

	"github.com/gorilla/mux"
)

func InitRoutes(server *server.Server) *mux.Router{
	router:=mux.NewRouter()

	utils.RouterUtils(router,"/v1",v1.V1(server))
	utils.RouterUtils(router,"/v2",v2.V2(server))
	
	return router
}