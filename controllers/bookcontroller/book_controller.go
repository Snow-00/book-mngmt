package bookcontroller

import (
  "net/http"
  "encoding/json"

  "github.com/Snow-00/book-mngmt/helper"
  "github.com/Snow-00/book-mngmt/entities"
  "github.com/Snow-00/book-mngmt/models/bookmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
  var books []entities.Book

  books, err := bookmodel.GetAll()
  if err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 200, "List of books", books)
}

func Create(w http.ResponseWriter, r *http.Request) {
  var book entities.Book

  if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }
  defer r.Body.Close()

  if err := bookmodel.Create(book); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 201, "Success create book", nil)
}