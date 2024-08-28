package api

import "net/http"

type Handler struct {
	Path        string
	Method      string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func NewHandler(path string, method string, handlerFunc func(w http.ResponseWriter, r *http.Request)) *Handler {
	return &Handler{
		Path:        path,
		Method:      method,
		HandlerFunc: handlerFunc,
	}
}
