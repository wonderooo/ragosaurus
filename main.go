package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"ragosaurus/pkg/api"
	"ragosaurus/pkg/api/handlers"
	"ragosaurus/pkg/util"
	"time"
)

func wfi() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	<-interruptChan
}

func main() {
	router := api.NewHttpRouter()
	router.RegisterHandlers(
		util.Merge(
			handlers.DocumentHandlers(),
		)...,
	)

	server := api.NewServer("0.0.0.0:8000", router)
	server.Run()

	wfi()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(server.ShutdownDeadline))
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)

	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	//
	//ctx := context.Background()
	//
	//client, err := chroma.NewClient("http://localhost:8750")
	//if err != nil {
	//	fmt.Printf("Failed to create client: %v", err)
	//}
	//
	//ef, err := openai.NewOpenAIEmbeddingFunction(os.Getenv("OPENAI_API_KEY"))
	//if err != nil {
	//	fmt.Printf("Failed to create client: %v", err)
	//}
	//
	//collection, err := client.CreateCollection(ctx, "my-collection", map[string]interface{}{"key1": "value1"}, true, ef, types.L2)
	//if err != nil {
	//	log.Fatalf("Failed to create collection: %v", err)
	//}
	//
	//_, err = collection.Add(context.TODO(), nil, []map[string]interface{}{{"key1": "value1"}}, []string{"My name is John and I have three dogs."}, []string{"ID1"})
	//if err != nil {
	//	log.Fatalf("Error adding documents: %v\n", err)
	//	return
	//}
}
