package main

import (
	"log"

	"github.com/machado-br/k8s-api/adapters/k8s"
	"github.com/machado-br/k8s-api/api"
	"github.com/machado-br/k8s-api/services/namespace"
)

func main() {

	adapter, err := k8s.NewAdapter(false)
	if err != nil {
		log.Fatalf("failed while creating k8s adapter: %v", err)
	}

	namespaceService := namespace.NewService(adapter)
	if err != nil {
		log.Fatalf("failed while creating k8s adapter: %v", err)
	}

	api, err := api.NewApi(adapter, namespaceService)
	if err != nil {
		log.Fatalf("failed while creating api: %v", err)
	}

	api.Run()
}
