package routes

import (
  "github.com/gorilla/mux"
  "github.com/Snow-00/book-mngmt/controllers/bookcontroller"
)

func BookRoutes(r *mux.Router){
  booksRoute := r.PathPrefix("/books").Subrouter()

  booksRoute.HandleFunc("", bookcontroller.Index).Methods("GET")
  booksRoute.HandleFunc("/create", bookcontroller.Create).Methods("POST")
  booksRoute.HandleFunc("/detail/{id}", bookcontroller.Detail).Methods("GET")
}