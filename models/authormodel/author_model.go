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

func Create(author entities.Author) error {
  _, err := config.DB.Exec(
    `INSERT INTO authors (name, gender, email, created_at, updated_at)
    VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);`,
    author.Name,
    author.Gender, 
    author.Email,
  )
  
  return err
}

func Detail(id uint) (entities.Author, error) {
  var author entities.Author

  row, err := config.DB.QueryRow(
    `SELECT * FROM authors WHERE id=?`,
    id,
  )

  if err != nil {
    return nil, err
  }
  defer row.Close()

  err := row.Scan(
    &author.ID,
    &author.Name,
    &author.Gender,
    &author.Email,
    &author.CreatedAt,
    &author.UpdatedAt,
  )

  return author, err
}