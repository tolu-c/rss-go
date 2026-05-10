package main

import (
	"log"
	"net/http"

	"github.com/tolu-c/rss-go/internal/api"
	"github.com/tolu-c/rss-go/internal/app"
	"github.com/tolu-c/rss-go/internal/pkg/environment"
	"github.com/tolu-c/rss-go/internal/store"
)

// main is intentionally short. Its job is to wire the dependency graph,
// nothing more — every line below should be either constructing a layer
// or starting the server.
func main() {
	env, err := environment.Load()
	if err != nil {
		log.Fatal(err)
	}

	s, err := store.New(env.DBURL)
	if err != nil {
		log.Fatalf("store: %v", err)
	}
	defer s.Close()

	dp := app.InitDp(env, s)

	apiHandler := api.New(dp)
	apiHandler.Build()

	server := &http.Server{
		Handler: apiHandler.Router(),
		Addr:    ":" + env.Port,
	}

	log.Printf("server starting at port %s", env.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
