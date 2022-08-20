package main

import (
	"log"
	"os"

	"github.com/machado-br/k8s-api/adapters/aws"
	"github.com/machado-br/k8s-api/adapters/k8s"
	"github.com/machado-br/k8s-api/services/createKubeConfig"
	"github.com/machado-br/k8s-api/services/describeCluster"
)

func main() {

	name := os.Getenv("CLUSTER_NAME")
	region := os.Getenv("AWS_REGION")

	cloudProviderAdapter, err := aws.NewAdapter(region, name)
	if err != nil {
		log.Fatalf("Failed while creating cloud provider adapter: %v", err)
	}

	describeClusterService, err := describeCluster.NewService(cloudProviderAdapter)
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
