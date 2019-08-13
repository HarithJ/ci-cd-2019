package auth

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetClientSet() *kubernetes.Clientset {
	// creates the in-cluster config
	config, err := rest.InClusterConfig() // returns pointer to config struct
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config) // returns pointer to clientset struct
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
