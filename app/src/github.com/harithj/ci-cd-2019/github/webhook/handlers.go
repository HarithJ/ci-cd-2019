package webhook

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"

  "github.com/harithj/ci-cd-2019/github/payload"
  "github.com/harithj/ci-cd-2019/pipeline/pipeline"
)

func Payloadfunc(w http.ResponseWriter, r *http.Request) {
  var err error
  var actualPayload payload.ActualPayload
  var payloadBytes []byte

  payloadBytes, err = ioutil.ReadAll(r.Body)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(payloadBytes)

  err = json.Unmarshal(payloadBytes, &actualPayload)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(actualPayload)

  payload := simplifyPayload(&actualPayload)

  pipeline.StartPipeline(payload)
}
