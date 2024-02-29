package signUp

import (
	"authenticationService/server"
	"net/http"
)

func SignUp(server *server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Hi i am sign Up"))
	}
}