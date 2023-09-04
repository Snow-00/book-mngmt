package bookmodel

import (
  "github.com/Snow-00/book-mngmt/entities"
  "github.com/Snow-00/book-mngmt/config"
)

func GetAll() ([]entities.Book, error) {
  var books []entities.Book

  rows, err := config.DB.Query(
    `SELECT
            b.id,
            b.title,
            a.name AS author,
            b.year,
            b.publisher,
            b.description,
            b.created_at,
            b.updated_at
     FROM books b
     JOIN authors a ON b.author_id = a.id;`,
  )

  if err != nil {
    return books, err
  }
  defer rows.Close()

  for rows.Next() {
    var book entities.Book
    err := rows.Scan(
      &book.ID,
      &book.Title,
      &book.Author.Name,
      &book.Year,
      &book.Publisher,
      &book.Description,
      &book.CreatedAt,
      &book.UpdatedAt,
    )

    if err != nil {
      return books, err
    }

    books = append(books, book)
  }

  return books, err
}