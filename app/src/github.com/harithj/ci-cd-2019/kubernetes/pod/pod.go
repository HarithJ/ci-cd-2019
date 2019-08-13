package pod

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

func CreatePod(clientset *kubernetes.Clientset, podYmlConfig string) {
	// create a pod client
	// *************************************************************************
	PodClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)
	//*************************************************************************

	// turn yml file to json
	// *************************************************************************
	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, _, err := decode([]byte(podYmlConfig), nil, nil)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	podConfig := obj.(*apiv1.Pod)
	//*************************************************************************

	// Create Pod
	// *************************************************************************
	fmt.Println("Creating pod...")
	result, err := PodClient.Create(podConfig)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	//*************************************************************************
}
