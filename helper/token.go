package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/Snow-00/book-mngmt/entities"
	"github.com/Snow-00/book-mngmt/config"
)

var expDate = 2 * time.Minute
var mySigningKey = []byte(config.ENV.SECRET)

type MyCustomClaims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.ResgisterdClaims
}

func CreateToken(user *entities.User) (string, error) {
	claims := MyCustomClaims{
		ID: user.ID,
		Username: user.Username,
		Role: user.Role,
		jwt.ResgisterdClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDate)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}