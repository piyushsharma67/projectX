package server

import "authenticationService/store"

type Server struct{
	Store *store.Store
}
func New(store *store.Store) *Server{
	server:=&Server{}

	server.Store=store

	return server
}