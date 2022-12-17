package main

import (
	"log"

	"github.com/machado-br/k8s-api/adapters/k8s"
	namespaceAdapter "github.com/machado-br/k8s-api/adapters/k8s/namespace"
	"github.com/machado-br/k8s-api/api"
	namespaceService "github.com/machado-br/k8s-api/services/namespace"
)

func main() {

	ClientSet, err := k8s.RetrieveClientSet(false)
	if err != nil {
		log.Fatalf("failed while creating k8s adapter: %v", err)
	}

	namespaceAdapter := namespaceAdapter.NewAdapter(ClientSet)
	if err != nil {
		log.Fatalf("failed while creating namespace adapter: %v", err)
	}

	namespaceService := namespaceService.NewService(namespaceAdapter)

	api, err := api.NewApi(namespaceService)
	if err != nil {
		log.Fatalf("failed while creating api: %v", err)
	}

	api.Run()
}
