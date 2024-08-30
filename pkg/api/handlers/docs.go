package handlers

import (
	"log"
	"net/http"
	"ragosaurus/pkg/api"
	"ragosaurus/pkg/rag"
)

func DocumentHandlers() []*api.Handler {
	return []*api.Handler{
		api.NewHandler("/api/document", "POST", addDocument),
	}
}

func addDocument(w http.ResponseWriter, r *http.Request) {
	err := rag.Client().ImportDocument("cipa3", "chuj", "1", nil)
	log.Println(err)

	w.Write([]byte("dupa"))
}
