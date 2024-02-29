package appRoutes

import (
	"authenticationService/server"
	v1 "authenticationService/server/controller/v1"
	utils "authenticationService/utils"

	"github.com/gorilla/mux"
)

func InitRoutes(server *server.Server) *mux.Router{
	router:=mux.NewRouter()

	utils.RouterUtils(router,"/v1",v1.V1())
	
	return router
}