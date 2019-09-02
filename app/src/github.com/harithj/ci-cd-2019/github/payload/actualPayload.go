package payload

type Repo struct {
	CloneUrl string `json:"clone_url"`
	Name     string `json:"name"`
}

type Head struct {
	Repo Repo   `json:"repo"`
  Ref  string `json:"ref"`
}

type PullRequest struct {
	Head Head `json:"head"`
}

type ActualPayload struct {
	PullRequest PullRequest `json:"pull_request"`
}
