package routes

import (
  "github.com/gorilla/mux"
  "github.com/Snow-00/book-mngmt/controllers/usercontroller"
)

func AuthRoutes(r *mux.Router) {
  authRoute := r.PathPrefix("/auth").Subrouter()

  authRoute.HandleFunc("/register", usercontroller.Register).Methods("POST")
  authRoute.HandleFunc("/login", usercontroller.Login).Methods("POST")
}