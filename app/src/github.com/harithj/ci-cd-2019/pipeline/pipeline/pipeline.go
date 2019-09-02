package pipeline

import (
  "io/ioutil"
  "fmt"

  "github.com/harithj/ci-cd-2019/github/payload"
  "github.com/harithj/ci-cd-2019/kubernetes/pod"
  "github.com/harithj/ci-cd-2019/kubernetes/auth"
	"github.com/harithj/ci-cd-2019/github/git"
)

func StartPipeline(payload *payload.Payload) {
  cloneDir := git.CloneRepo(payload)

  podYmlConfig, err := ioutil.ReadFile(cloneDir + "/2019/pod.yml")
	if err != nil {
		fmt.Println(err)
	}

  clientset := auth.GetClientSet()
	pod.CreatePod(clientset, string(podYmlConfig))
}
