package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Snow-00/book-mngmt/config"
  "github.com/Snow-00/book-mngmt/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

  r := mux.NewRouter()
  routes.IndexRoute(r)

	log.Println("Server is running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
