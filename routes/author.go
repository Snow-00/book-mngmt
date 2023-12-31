package routes

import (
  "github.com/gorilla/mux"
  "github.com/Snow-00/book-mngmt/controllers/authorcontroller"
)

func AuthorRoutes(r *mux.Router){
  authorRoute := r.PathPrefix("/authors").Subrouter()

  authorRoute.HandleFunc("", authorcontroller.Index).Methods("GET")
  authorRoute.HandleFunc("/create", authorcontroller.Create).Methods("POST")
  authorRoute.HandleFunc("/detail/{id}", authorcontroller.Detail).Methods("GET")
  authorRoute.HandleFunc("/update/{id}", authorcontroller.Update).Methods("PUT")
  authorRoute.HandleFunc("/delete/{id}", authorcontroller.Delete).Methods("DELETE")
}