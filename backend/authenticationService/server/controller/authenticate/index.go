package authenticate

import (
	"authenticationService/server"
	v1 "authenticationService/server/controller/authenticate/v1"
	"authenticationService/utils"

	"github.com/gorilla/mux"
)

func AuthenticateInitRoute(server *server.Server)*mux.Router{
	router:=mux.NewRouter()

	utils.RouterUtils(router,"/v1",v1.V1(server))

	return router
}