package usercontroller

import (
	"encoding/json"
  "database/sql"
	"net/http"
  "errors"

	"github.com/Snow-00/book-mngmt/helper"
  "github.com/Snow-00/book-mngmt/entities"
	"github.com/Snow-00/book-mngmt/models/usermodel"
)

func Register(w http.ResponseWriter, r *http.Request) {
  var register entities.Register

  if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }
  defer r.Body.Close()

  if register.Password != register.ConfirmPassword {
    helper.Response(w, 400, "Password not match", nil)
    return
  }

  // create hash password
  hashPassword, err := helper.HashPassword(register.Password)
  if err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  user := entities.User{
    Username: register.Username,
    HashPassword: hashPassword,
    Role: register.Role,
  }
  
  if err := usermodel.Register(user); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 201, "Register success", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
  var login entities.Login

  if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }
  defer r.Body.Close()

  // check if the username exists
  user, err := usermodel.Login(login.Username)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      helper.Response(w, 404, "User not found", nil)
      return
    }

    helper.Response(w, 500, err.Error(), nil)
    return
  }

  // verify the password
  if err := helper.VerifyPassword(user.HashPassword, login.Password); err != nil {
    helper.Response(w, 404, "Wrong Password", nil)
    return
  }

  // create token
  token, errToken := helper.CreateToken(&user)
  if errToken != nil {
    helper.Response(w, 500, errToken.Error(), nil)
    return
  }

  helper.Response(w, 200, "Success Login", token)
}