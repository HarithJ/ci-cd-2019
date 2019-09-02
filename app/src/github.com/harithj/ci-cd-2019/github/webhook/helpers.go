package webhook

import (
  "github.com/harithj/ci-cd-2019/github/payload"
)

func simplifyPayload(actualPayload *payload.ActualPayload) *payload.Payload{
  return &payload.Payload {
    RepoName: actualPayload.PullRequest.Head.Repo.Name,
    CloneUrl: actualPayload.PullRequest.Head.Repo.CloneUrl,
    Ref:      actualPayload.PullRequest.Head.Ref,
  }
}
