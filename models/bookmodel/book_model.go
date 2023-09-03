package bookmodel

import (
  "github.com/Snow-00/book-mngmt/entities"
  "github.com/Snow-00/book-mngmt/config"
)

func GetAll() ([]entities.Book, error) {
  var books []entites.Book

  rows, err := config.DB.Exec(`SELECT * FROM books;`)
  if err != nil {
    return books, err
  }
  defer rows.Close()

  for rows.Next() {
    var book entities.Book
    err := rows.Scan(
      &book.ID,
      &book.Title,
      &book.AuthorID,
      &book.Year,
      &book.Publication,
      &book.Description,
      &book.CreatedAt,
      &book.UpdatedAt,
    )

    if err != nil {
      return books, err
    }

    books = append(books, book)
  }

  return books, book
}