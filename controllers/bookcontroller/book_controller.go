package bookcontroller

import (
  "net/http"

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