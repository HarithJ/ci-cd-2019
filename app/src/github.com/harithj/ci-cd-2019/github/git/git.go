package git

import (
  "os"
  "fmt"

  "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/plumbing"

  "github.com/harithj/ci-cd-2019/github/payload"
)

func CloneRepo(payload *payload.Payload) string {
  cloneDir := "/temp/" + payload.RepoName
  os.Mkdir(cloneDir, 0777)

  _, err := git.PlainClone(cloneDir, false, &git.CloneOptions{
    URL: payload.CloneUrl,
    ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", payload.Ref)),
  })

  if err != nil {
    fmt.Println(err)
  }

  return cloneDir
}
