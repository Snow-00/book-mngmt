package routes

import (
  "github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
  userRoute := r.PathPrefix("").Subrouter()
}