package authormodel

import (
  "github.com/Snow-00/book-mngmt/entities"
  "github.com/Snow-00/book-mngmt/config"
)
func GetAll() any {
  rows, err := config.DB.Query(`SELECT * FROM authors`)
  if err != nil {
    rows = []entities.Author
  }
}