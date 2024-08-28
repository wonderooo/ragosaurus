package handlers

import (
	"net/http"
	"ragosaurus/pkg/api"
)

func DocumentHandlers() []*api.Handler {
	return []*api.Handler{
		api.NewHandler("/api/document", "POST", addDocument),
	}
}

func addDocument(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("dupa"))
}
