package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/Snow-00/my-gobookstore/pkg/routes"
)

func main() {
  r := mux.NewRouter()
  routes.RegisterBookStoreRoutes(r)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(":8080", r))
}