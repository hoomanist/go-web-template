package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (app *App) HandleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{
				"error": "not parsable request",
			})

			w.Write(res)
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(r.PostForm.Get("password")), bcrypt.DefaultCost)
		hasher := md5.New()
		hasher.Write(hash)
		token := hex.EncodeToString(hasher.Sum(nil))
		err = app.db.Create(&User{
			Username: r.PostForm.Get("username"),
			Password: r.PostForm.Get("password"),
			Email:    r.PostForm.Get("email"),
			Token:    token,
		}).Error
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{
				"error": "cannot insert to db",
			})
			w.Write(res)
		}

		res, _ := json.Marshal(map[string]interface{}{
			"token": token,
		})
		w.Write(res)
	}
}

func (app *App) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{
				"error": "not parsable request",
			})

			w.Write(res)
		}
		var user User
		err = app.db.Where(&User{
			Username: r.PostForm.Get("username"),
			Password: r.PostForm.Get("password"),
		}).First(&user).Error
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{
				"error": "there is no such user.",
			})
			w.Write(res)
		}
		res, _ := json.Marshal(map[string]interface{}{
			"token": user.Token,
		})
		w.Write(res)
	}
}
