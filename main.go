package main

import (
	"github.com/Snow-00/book-mngmt/config"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()

	log.Println("Server is running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
