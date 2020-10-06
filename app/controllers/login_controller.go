package controllers

import (
	"Rest/app/auth"
	"Rest/app/models"
	u "Rest/app/utils"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)
/**
Авторизация пользователя с последующим входом и генерацией токена
 */
func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := a.SignIn(user.Email, user.Password)
	if err != nil {

		formattedError := u.KindError(err.Error())
		u.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	u.JSON(w, http.StatusOK, token)
}
/**
Здесь проверка на существование пользователя в базе данных и проверка пароля

*/
func (a *App) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = a.DB.QueryRow("select * from users where email=$1 and password=$2", email,password).Scan(&user.ID, &user.Nickname, &user.Email, &user.Password)
	if err != nil {

		return "", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {

		return "", err
	}


	return auth.CreateToken(user.ID)
}

