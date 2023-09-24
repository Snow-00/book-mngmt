package token

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/Snow-00/book-mngmt/entities"
	"github.com/Snow-00/book-mngmt/config"
)

type MyCustomClaims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.ResgisterdClaims
}

func CreateToken(user *entities.User) (string, error) {
	claims := 	
}