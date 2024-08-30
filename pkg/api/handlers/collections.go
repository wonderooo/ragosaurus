package handlers

import (
	"net/http"
	"ragosaurus/pkg/api"
	"ragosaurus/pkg/rag"
)

func CollectionHandlers() []*api.Handler {
	return []*api.Handler{
		api.NewHandler("/api/collection", "POST", addCollection),
	}
}

func addCollection(w http.ResponseWriter, r *http.Request) {
	rag.Client().CreateCollection("cipa")
	w.Write([]byte("dupa2"))
}
