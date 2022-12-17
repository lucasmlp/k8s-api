package main

import (
	"log"

	"github.com/machado-br/k8s-api/api"
)

func main() {
	api, err := api.NewApi()
	if err != nil {
		log.Fatalf("failed while creating api: %v", err)
	}

	api.Run()
}
