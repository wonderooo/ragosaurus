package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"ragosaurus/pkg/api"
	"ragosaurus/pkg/api/handlers"
	"ragosaurus/pkg/rag"
	"ragosaurus/pkg/util"
	"time"
)

func wfi() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	<-interruptChan
}

func env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	env()

	rag.Init()

	router := api.NewHttpRouter()
	router.RegisterHandlers(
		util.Merge(
			handlers.DocumentHandlers(),
			handlers.CollectionHandlers(),
		),
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
}
