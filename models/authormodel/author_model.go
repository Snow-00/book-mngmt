package authormodel

import (
  "github.com/Snow-00/book-mngmt/entities"
  "github.com/Snow-00/book-mngmt/config"
)

func GetAll() ([]entities.Author, error) {
  var authors []entities.Author

  rows, err := config.DB.Query(`SELECT * FROM authors`)
  if err != nil {
    return authors, err
  }
  defer rows.Close()

  for rows.Next() {
    var author entities.Author
    err := rows.Scan(
      &author.ID,
      &author.Name,
      &author.Gender,
      &author.Email,
      &author.CreatedAt,
      &author.UpdatedAt,
    )
    if err != nil {
      return authors, err
    }

    authors = append(authors, author)
  }

  return authors, err
}