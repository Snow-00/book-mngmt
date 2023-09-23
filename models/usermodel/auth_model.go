package usermodel

import (
  "github.com/Snow-00/book-mngmt/config"
  "github.com/Snow-00/book-mngmt/entities"
)

func Register(user entities.User) error {
  _, err := config.DB.Exec(
    `INSERT INTO users (username, hash_password, role, is_deleted, created_at, updated_at)
    VALUES (?, ?, ?, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
    user.Username, user.HashPassword, user.Role,
  )

  return err
}

func Login(username string) (entities.User, error) {
  var user entities.User

  row := config.DB.QueryRow(
    `SELECT * FROM users WHERE username = ? AND is_deleted = 0`,
    username,
  )

  err := row.Scan(
    &user.ID,
    &user.Username,
    &user.HashPassword,
    &user.Role,
    &user.IsDeleted,
    &user.CreatedAt,
    &user.UpdatedAt,
  )

  return user, err
}