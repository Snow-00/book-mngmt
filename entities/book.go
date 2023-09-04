package entities

import "time"

type Book struct {
  ID uint `json:"id"`
  Title string `json:"title"`
  Author Author `json:"author"`
  Year int `json:"year"`
  Publisher string `json:"publisher"`
  Description string `json:"description"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}