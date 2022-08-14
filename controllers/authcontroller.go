package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/hioan-dev/go-auth/entities"
	"github.com/hioan-dev/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

var UserModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {

	temp, _ := template.ParseFiles("views/index.html")

	temp.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// Proses login
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user entities.User
		UserModel.Where(&user, "username", UserInput.Username)

		var message error
		if user.Username == "" {
			message = errors.New("Username atau password salah")
		} else {
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			if errPassword != nil {
				message = errors.New("Username atau password salah")
			}
		}

		if message != nil {

			data := map[string]interface{}{
				"message": message.Error(),
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		}
	}
}
