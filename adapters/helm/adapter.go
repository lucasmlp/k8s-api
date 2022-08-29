package helm

import (
	"log"
	"os"

	"github.com/machado-br/k8s-api/adapters/models"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
)

type adapter struct{
	action *action.Configuration
}

type Adapter interface{
	ListReleases() []models.Release
}

func NewAdapter(
) (adapter, error) {
	kubeconfigPath := "./config/kube"
	releaseNamespace := "default"

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(kubeconfigPath, "", releaseNamespace), releaseNamespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		panic(err)
	}

    // settings := cli.New()
	// actionConfig := new(action.Configuration)
	// if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
    //     log.Printf("%+v", err)
    //     os.Exit(1)
    // }

	return adapter{
		action: actionConfig,
	}, nil
}
func (a adapter) ListReleases() []models.Release{

	listAction := action.NewList(a.action)
	releases, err := listAction.Run();
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("releases: %v\n", releases)
	return []models.Release{}
}
