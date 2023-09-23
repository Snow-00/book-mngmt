package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Snow-00/book-mngmt/helper"
  "github.com/Snow-00/book-mngmt/entities"
	"github.com/Snow-00/book-mngmt/usermodel"
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

  hashPassword, err := helper.HashPassword(register.Password)
  if err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  user := usermodel.User{
    Username: register.Username,
    HashPassword: register.HashPassword,
    Role: register.Role,
  }
  
  if err := usermodel.Register(user); err != nil {
    helper.Response(w, 500, err.Error(), nil)
    return
  }

  helper.Response(w, 201, "Register success", nil)
}