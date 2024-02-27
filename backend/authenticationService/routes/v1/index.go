package v1

import (
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

	router.Handle("signUp",signUp)

	return router
}