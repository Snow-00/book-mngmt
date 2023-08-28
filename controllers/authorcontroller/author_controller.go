package authorcontroller

import (
  "net/http"

  "github.com/Snow-00/book-mngmt/entities"
  "github.com/Snow-00/book-mngmt/config"
)

func Index(w http.ResponseWriter, r *http.Request){
  var authors []entities.Author

  if err:= config.DB
}