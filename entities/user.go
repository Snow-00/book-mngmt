package entities

import "time"

type User struct {
  ID int `json:"id"`
  Username string `json:"username"`
  HashPassword string `json:"hash_password"`
  Role string `json:"role"`
  IsDeleted bool `json:"is_deleted"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

type Register struct {
  Username string `json:"username"`
  Password string `json:"password"`
  ConfirmPassword string `json:"confirm_password"`
  Role string `json:"role"`
}

type Login struct {
  Username string `json:"username"`
  Password string `json:"password"` 
  Role string `json:"role"`
}