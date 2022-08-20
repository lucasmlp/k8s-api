package main

import (
	"log"

	"github.com/machado-br/k8s-api/adapters/aws"
	"github.com/machado-br/k8s-api/adapters/k8s"
	"github.com/machado-br/k8s-api/services/createKubeConfig"
	describecluster "github.com/machado-br/k8s-api/services/describeCluster"
)

const (
	name = "e-commerce"
    region = "us-west-2"
)

func main() {

	cloudProviderAdapter, err := aws.NewAdapter(region, name)
	if err != nil {
		log.Fatalf("Failed while creating cloud provider adapter: %v", err)
	}

	describeClusterService, err := describecluster.NewService(cloudProviderAdapter)
	if err != nil {
		log.Fatalf("Failed while creating createKubeConfig service: %v", err)
	}

	cluster, err := describeClusterService.Run()
	if err != nil {
		log.Fatalf("Failed while retrieving cluster information: %v", err)
	}

	k8sAdapter, err := k8s.NewAdapter(cluster)
	if err != nil {
		log.Fatalf("Failed while creating k8s adapter: %v", err)
	}

	createKubeConfigService, err := createKubeConfig.NewService(name, region, k8sAdapter)
	if err != nil {
		log.Fatalf("Failed while creating createKubeConfig service: %v", err)
	}

	err = createKubeConfigService.Run()
	if err != nil {
		log.Fatalf("Failed while creating createKubeConfig file: %v", err)
	}

	log.Println("Kubeconfig file created successfuly")
}
