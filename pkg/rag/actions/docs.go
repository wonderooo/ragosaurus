package actions

import (
	"context"
	chroma "github.com/amikos-tech/chroma-go"
	"github.com/amikos-tech/chroma-go/openai"
)

func ImportDocument(
	chroma *chroma.Client,
	embeddingFunction *openai.OpenAIEmbeddingFunction,
	collectionName string,
	content string,
	id string,
	metadata map[string]interface{},
) error {
	ctx := context.TODO()

	collection, err := GetCollection(chroma, embeddingFunction, collectionName)
	if err != nil {
		return err
	}

	_, err = collection.Add(ctx, nil, []map[string]interface{}{metadata}, []string{content}, []string{id})

	return err
}
