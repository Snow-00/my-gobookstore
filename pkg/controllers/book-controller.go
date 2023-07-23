package controllers

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "strconv"
  "github.com/Snow-00/my-gobookstore/utils"
  "github.com/Snow-00/my-gobookstore/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
  newBooks := models.GetAllBooks()
  res, _ := json.Marshal(newBooks)
}