package main

import (
	"github.com/harithj/ci-cd-2019/auth"
	"github.com/harithj/ci-cd-2019/pod"
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

func main() {
	clientset := auth.GetClientSet()
	pod.CreatePod(clientset, json)
}
