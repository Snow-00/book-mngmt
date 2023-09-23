package usermodel

import (
  "github.com/Snow-00/book-mngmt/config"
  "github.com/Snow-00/book-mngmt/entities"
)

func Register(user entities.User) error {
  _, err := config.DB.Exec(
    `INSERT INTO users (username, hash_password, role, is_deleted, created_at, updated_at)
    VALUES (?, ?, ?, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
    user.Username, user.HashPassword, user.Role, user.IsDeleted,
  )

  return err
}