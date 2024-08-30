package actions

import (
	"context"
	"errors"
	chroma "github.com/amikos-tech/chroma-go"
	"github.com/amikos-tech/chroma-go/openai"
	"github.com/amikos-tech/chroma-go/types"
)

func CreateCollection(
	chroma *chroma.Client,
	embeddingFunction *openai.OpenAIEmbeddingFunction,
	name string,
) error {
	ctx := context.TODO()
	_, err := chroma.CreateCollection(ctx, name, nil, true, embeddingFunction, types.L2)
	return errors.Join(errors.New("chroma create collection error"), err)
}

func GetCollection(
	chroma *chroma.Client,
	embeddingFunction *openai.OpenAIEmbeddingFunction,
	name string,
) (*chroma.Collection, error) {
	ctx := context.TODO()
	collection, err := chroma.GetCollection(ctx, name, embeddingFunction)
	err = errors.Join(errors.New("chroma get collection error"), err)
	return collection, err
}
