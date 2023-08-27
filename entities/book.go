package entities

import "time"

type Book struct {
  ID uint `json:"id"`
  Title string `json:"title"`
  AuthorID uint `json:"author_id"`
  Year int `json:"year"`
  Publisher string `json:"publisher"`
  Description string `json:"description"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}