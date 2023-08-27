package entities

import "time"

type Author struct {
  ID uint `json:"id"`
  Name string `json:"name"`
  Gender string `json:"gender"`
  Email string `json:"email"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}