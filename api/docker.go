package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
)

type DockerHandler struct {
	dockerClient *client.Client
}

func NewDockerHandler() (*DockerHandler, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &DockerHandler{dockerClient: cli}, nil
}

func (h *DockerHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/containers", h.ListContainers)
	return r
}

func (h *DockerHandler) ListContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := h.dockerClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, container := range containers {
		fmt.Fprintf(w, "ID: %s, Image: %s, Status: %s\n", container.ID[:10], container.Image, container.Status)
	}
}
