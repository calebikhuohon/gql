package main

import (
	"context"
	"fmt"
	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"quicknode/graphql/client"
	"quicknode/graphql/server"
	"quicknode/graphql/server/generated"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		systemCall := <-exit
		log.Printf("System call: %+v", systemCall)
		cancel()
	}()

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Failed to parse app config: %s", err)
	}
	if err := config.Validate(); err != nil {
		log.Fatalf("Failed to validate config: %s", err)
	}

	cli := client.NewClient(config.ApiKey)
	gqlRouter := chi.NewRouter()
	gqlRouter.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            false,
	}).Handler)

	srv := gqlHandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &server.Resolver{
			Client: cli,
		},
	}))
	srv.Use(extension.Introspection{})
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	gqlRouter.Handle("/query", srv)
	gqlRouter.Handle("/query/playground", playground.Handler("GraphQL playground", "/query"))
	go func() {
		log.Print("starting graphql server...")
		if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Ports.GRAPHQL), gqlRouter); err != nil {
			log.Printf("Could not listen and serve: %s", err)
			os.Exit(1)
		}
	}()
	<-ctx.Done()

	log.Print("Stopping the graphql server..")
}
