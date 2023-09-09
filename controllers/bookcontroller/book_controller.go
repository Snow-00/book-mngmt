package bookcontroller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Snow-00/book-mngmt/entities"
	"github.com/Snow-00/book-mngmt/helper"
	"github.com/Snow-00/book-mngmt/models/bookmodel"
	"github.com/gorilla/mux"
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

func Detail(w http.ResponseWriter, r *http.Request) {
  idParam := mux.Vars(r)["id"]
  id, _ := strconv.Atoi(idParam)

  book, err := bookmodel.Detail(id)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      helper.Response(w, 404, "Book not found", nil)
      return
    }
    
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 200, "Book Detail", book)
}

func Update(w http.ResponseWriter, r *http.Request) {
  var book entities.Book
  
  idParam := mux.Vars(r)["id"]
  id, _ := strconv.Atoi(idParam)

  // check data exist
  _, err := bookmodel.Detail(id)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      helper.Response(w, 404, "Book not found", nil)
      return
    }

    helper.Response(w, 500, err.Error(), nil)
    return
  }

  if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  if err := bookmodel.Update(book, id); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 201, "Update book success", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
  idParam := mux.Vars(r)["id"]
  id, _ := strconv.Atoi(idParam)

  _, err := bookmodel.Detail(id)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      helper.Response(w, 404, "Book not found", nil)
      return
    }

    helper.Response(w, 500, err.Error(), nil)
    return
  }

  if err := bookmodel.Delete(id); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 200, "Success delete book", nil)
}