package v1

import (
	"authenticationService/routes/v1/signUp"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func mount(r *mux.Router, path string, handler http.Handler) {
    r.PathPrefix(path).Handler(
        http.StripPrefix(
            strings.TrimSuffix(path, "/"),
            handler,
        ),
    )
}

func v1Routes() *mux.Router{
	router:=mux.NewRouter()

	router.HandleFunc("signUp",signUp.SignUp)

	return router
}