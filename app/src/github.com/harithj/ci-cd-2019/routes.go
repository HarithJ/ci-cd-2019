package main

import (
    "net/http"

    "github.com/gorilla/mux"

    "github.com/harithj/ci-cd-2019/github/webhook"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  for _, route := range routes {
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(route.HandlerFunc)
  }

  return router
}

var routes = Routes{
  Route{
    "Index",
    "Get",
    "/",
    Index,
  },
  Route{
    "Payload",
    "POST",
    "/payload",
    webhook.Payloadfunc,
  },
}
