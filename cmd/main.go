package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/machado-br/k8s-api/adapters/k8s"
	namespaceAdapter "github.com/machado-br/k8s-api/adapters/k8s/namespace"
	"github.com/machado-br/k8s-api/api"
	namespaceService "github.com/machado-br/k8s-api/services/namespace"
)

func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("failed to retrieve string variable %v", key))
	}
	return val
}

func getBoolEnv(key string) bool {
	stringValue := getStrEnv(key)
	ret, err := strconv.ParseBool(stringValue)
	if err != nil {
		panic(fmt.Sprintf("failed to retrieve boolean variable %v", key))
	}

	return ret
}

func main() {
	deployed := getBoolEnv("DEPLOYED")

	ClientSet, err := k8s.RetrieveClientSet(deployed)
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
