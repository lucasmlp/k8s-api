package k8s

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type adapter struct {
	Deployed  bool
	ClientSet *kubernetes.Clientset
}

func NewAdapter(deployed bool) (adapter, error) {
	clientSet, err := retrieveClientSet(deployed)
	if err != nil {
		return adapter{}, err
	}

	return adapter{
		Deployed:  deployed,
		ClientSet: clientSet,
	}, nil
}

func retrieveClientSet(deployed bool) (*kubernetes.Clientset, error) {
	var clientset *kubernetes.Clientset

	if deployed {
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}

		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
	} else {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
	}

	return clientset, nil
}
