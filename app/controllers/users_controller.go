package controllers

import (
	"Rest/app/auth"
	"Rest/app/models"
	u "Rest/app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.CreateUser(a.DB)

	if err != nil {

		formattedError := u.KindError(err.Error())

		u.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	u.JSON(w, http.StatusCreated, userCreated)
}

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllUsers(a.DB)
	if err != nil {
		u.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	u.JSON(w, http.StatusOK, users)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	var err error
	vars := mux.Vars(r)
	user.ID, err = strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_,err = user.FindUserByID(a.DB)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}
	u.JSON(w, http.StatusOK, user)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}
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

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		u.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != int(id) {
		u.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.ID = id
	updated,err := user.UpdateAUser(a.DB)
	if err != nil {
		formattedError := u.KindError(err.Error())
		u.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	u.JSON(w, http.StatusOK, updated)
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.User{}
	var err error
	user.ID, err = strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		u.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != 0 && uint64(tokenID) != user.ID {
		u.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	_,err = user.DeleteAUser(a.DB)
	if err != nil {
		u.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", user.ID))
	u.JSON(w, http.StatusNoContent, "")
}
