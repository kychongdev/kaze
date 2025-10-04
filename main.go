package main

import (
	"net/http"

	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kychongdev/kaze/api"
)

func main() {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dockerHandler, err := api.NewDockerHandler()
	if err != nil {
		panic(err)
	}
	r.Mount("/api/docker", dockerHandler.Routes())

	http.ListenAndServe(":3001", r)
}
