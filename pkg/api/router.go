package api

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Inner *mux.Router
}

func NewHttpRouter() *Router {
	return &Router{
		Inner: mux.NewRouter(),
	}
}

func (router *Router) RegisterHandlers(handlers ...*Handler) {
	for _, handler := range handlers {
		router.Inner.
			HandleFunc(handler.Path, handler.HandlerFunc).
			Methods(handler.Method).
			Schemes("http", "https")
	}
}
