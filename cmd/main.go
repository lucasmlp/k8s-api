package main

import (
	"log"
	"os"

	"github.com/machado-br/k8s-api/adapters/aws"
	"github.com/machado-br/k8s-api/adapters/helm"
	"github.com/machado-br/k8s-api/adapters/k8s"
	"github.com/machado-br/k8s-api/api"
	"github.com/machado-br/k8s-api/services/createKubeConfig"
	"github.com/machado-br/k8s-api/services/describeCluster"
	"github.com/machado-br/k8s-api/services/listReleases"
)

func main() {

	name := os.Getenv("CLUSTER_NAME")
	region := os.Getenv("AWS_REGION")
	namespace := os.Getenv("NAMESPACE")

	awsAdapter, err := aws.NewAdapter(region, name)
	if err != nil {
		log.Fatalf("failed while creating cloud provider adapter: %v", err)
	}

	describeClusterService, err := describeCluster.NewService(awsAdapter)
	if err != nil {
		log.Fatalf("failed while creating createKubeConfig service: %v", err)
	}

	cluster, err := describeClusterService.Run()
	if err != nil {
		log.Fatalf("failed while retrieving cluster information: %v", err)
	}

	k8sAdapter, err := k8s.NewAdapter(cluster, namespace, region)
	if err != nil {
		log.Fatalf("failed while creating k8s adapter: %v", err)
	}

	createKubeConfigService, err := createKubeConfig.NewService(k8sAdapter)
	if err != nil {
		log.Fatalf("failed while creating createKubeConfig service: %v", err)
	}

	err = createKubeConfigService.Run()
	if err != nil {
		log.Fatalf("failed while creating createKubeConfig file: %v", err)
	}

	log.Println("Kubernetes configuration file created successfully")

	kubeconfigPath := "./config/kube"

	helmDriver := os.Getenv("HELM_DRIVER")

	helmAdapter, err := helm.NewAdapter(namespace, kubeconfigPath, helmDriver)
	if err != nil {
		log.Fatalf("failed while creating helm adapter: %v", err)
	}

	listReleasesService, err := listReleases.NewService(helmAdapter)
	if err != nil {
		log.Fatalf("failed while creating list releases service: %v", err)
	}

	api, err := api.NewApi(listReleasesService)
	if err != nil {
		log.Fatalf("failed while creating api: %v", err)
	}

	api.Run()
}
