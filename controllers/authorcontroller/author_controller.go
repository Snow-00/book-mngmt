package authorcontroller

import (
  "net/http"

  // "github.com/Snow-00/book-mngmt/entities"
  // "github.com/Snow-00/book-mngmt/config"
  "github.com/Snow-00/book-mngmt/models/authormodel"
  "github.com/Snow-00/book-mngmt/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
  authors, err := authormodel.GetAll()

  if err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 200, "List of Authors", authors)
}

func Create(w http.ResponseWriter, r *http.Request) {
  
}