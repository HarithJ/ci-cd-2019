package main

import (
  "fmt"
  "net/http"

  "github.com/harithj/ci-cd-2019/kubernetes/auth"
  "github.com/harithj/ci-cd-2019/kubernetes/pod"
)

var json = `
apiVersion: v1
kind: Pod
metadata:
  name: pod-created
spec:
  containers:
    - name: nginx
      image: nginx:1.7.9
      ports:
      - containerPort: 80
`

func Index(w  http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Head to /create-pod to create pod")
}

func CreatePod(w  http.ResponseWriter, r *http.Request) {
  clientset := auth.GetClientSet()
	pod.CreatePod(clientset, json)
}
