package rag

import (
	"context"
	"github.com/amikos-tech/chroma-go"
	"github.com/amikos-tech/chroma-go/openai"
	"log"
	"os"
	"ragosaurus/pkg/rag/actions"
	"sync"
)

var lock = sync.Mutex{}
var rag *Rag

type Rag struct {
	chroma            *chromago.Client
	embeddingFunction *openai.OpenAIEmbeddingFunction
}

func Client() *Rag {
	if rag == nil {
		tryCreateInstance()
	}

	return rag
}

func Init() {
	if rag == nil {
		tryCreateInstance()
	}
}

func tryCreateInstance() {
	lock.Lock()
	defer lock.Unlock()

	if rag == nil {
		rag = newRag()
	}
}

func newRag() *Rag {
	chroma, err := chromago.NewClient("http://localhost:8750")
	if err != nil {
		log.Fatal("Error: Creating ChromaDb client")
	}
	_, err = chroma.Heartbeat(context.TODO())
	if err != nil {
		log.Fatal("Error: Heartbeat failed")
	}
	log.Println("Chroma client initialized")

	openaiKey := os.Getenv("OPENAI_API_KEY")
	embeddingFunction, err := openai.NewOpenAIEmbeddingFunction(openaiKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created embedding function")

	return &Rag{
		chroma:            chroma,
		embeddingFunction: embeddingFunction,
	}
}

func (rag *Rag) ImportDocument(collection string, content string, id string, metadata map[string]interface{}) error {
	return actions.ImportDocument(rag.chroma, rag.embeddingFunction, collection, content, id, metadata)
}

func (rag *Rag) CreateCollection(name string) error {
	return actions.CreateCollection(rag.chroma, rag.embeddingFunction, name)
}
