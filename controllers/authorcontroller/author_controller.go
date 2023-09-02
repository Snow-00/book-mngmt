package authorcontroller

import (
  "strconv"
  "errors"
  "net/http"
  "encoding/json"
  "database/sql"

  "github.com/Snow-00/book-mngmt/entities"
  // "github.com/Snow-00/book-mngmt/config"
  "github.com/Snow-00/book-mngmt/models/authormodel"
  "github.com/Snow-00/book-mngmt/helper"
  "github.com/gorilla/mux"
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
  var author entities.Author

  if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }
  defer r.Body.Close()
  
  if err := authormodel.Create(author); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 201, "Success create author", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
  idParam := mux.Vars(r)["id"]
  id, _ := strconv.Atoi(idParam)

  author, err := authormodel.Detail(id)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      helper.Response(w, 404, "Author not found", nil)
      return
    }

    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 200, "Detail of Author", author)
}

func Update(w http.ResponseWriter, r *http.Request) {
  var author entities.Author
  
  idParam := mux.Vars(r)["id"]
  id, _ := strconv.Atoi(idParam)

  // check whether data exist
  _, err := authormodel.Detail(id)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      helper.Response(w, 404, "Author not found", nil)
      return
    }

    helper.Response(w, 500, err.Error(), nil)
    return
  }

  // update data
  if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }
  defer r.Body.Close()

  if err := authormodel.Update(author, id); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 201, "Success update author", nil)
}