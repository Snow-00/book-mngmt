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
            b.year,
            b.publisher,
            b.description,
            b.created_at,
            b.updated_at,
            a.id AS author_id,
            a.name AS author
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
      &book.Year,
      &book.Publisher,
      &book.Description,
      &book.CreatedAt,
      &book.UpdatedAt,
      &book.Author.ID,
      &book.Author.Name,
    )

    if err != nil {
      return books, err
    }

    books = append(books, book)
  }

  return books, err
}

func Create(book entities.Book) error {
  _, err := config.DB.Exec(
    `INSERT INTO books (title, year, publisher, description, author_id, created_at, updated_at)
    VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
    &book.Title,
    &book.Year,
    &book.Publisher,
    &book.Description,
    &book.Author.ID,
  )

  return err
}

func Detail(id int) (entities.Book, error) {
  var book entities.Book

  row := config.DB.QueryRow(
    `SELECT
            b.id,
            b.title,
            b.year,
            b.publisher,
            b.description,
            b.created_at,
            b.updated_at,
            a.id AS author_id,
            a.name AS author
    FROM books b
    JOIN authors a ON b.author_id = a.id
    WHERE b.id = ?`,
    id,
  )

  err := row.Scan(
    &book.ID,
    &book.Title,
    &book.Year,
    &book.Publisher,
    &book.Description,
    &book.CreatedAt,
    &book.UpdatedAt,
    &book.Author.ID,
    &book.Author.Name,
  )

  return book, err
}

func Update(book entities.Book, id int) error {
  _, err := config.DB.Exec(
    `UPDATE books
    SET title = ?, year = ?, publisher = ?, description = ?, author_id = ?, updated_at = CURRENT_TIMESTAMP
    WHERE id = ?`,
    &book.Title,
    &book.Year,
    &book.Publisher,
    &book.Description,
    &book.Author.ID,
    id,
  )

  return err
}

func Delete(id int) error {
  _, err := config.DB.Exec(
    `DELETE FROM books WHERE id = ?`,
    id,
  )

  return err
}