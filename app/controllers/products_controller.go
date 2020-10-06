package controllers

import (
	"Rest/app/auth"
	"Rest/app/models"
	u "Rest/app/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a *App) GetProducts(w http.ResponseWriter, req *http.Request) {
	var p models.Product
	arr, err := p.GetProducts(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			u.RespondWithError(w, http.StatusNotFound, "Product isn't found")
		default:
			u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	u.RespondWithJSON(w, http.StatusOK, arr)

}

func (a *App) GetProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
	}
	p := models.Product{ID: id}
	_,err = p.GetProduct(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			u.RespondWithError(w, http.StatusNotFound, "Product isn't found")
		default:
			u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	u.RespondWithJSON(w, http.StatusOK, p)
}

func (a *App) CreateProduct(w http.ResponseWriter, req *http.Request) {

	var p models.Product
	decoder := json.NewDecoder(req.Body)
	uid, err := auth.ExtractTokenID(req)
	if err != nil {
		u.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if err := decoder.Decode(&p); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	p.UserID = uid
	p.Prepare()
	err = p.Validate()

	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if _,err := p.CreateProduct(a.DB); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	u.RespondWithJSON(w, http.StatusCreated, p)

}

func (a *App) DeleteProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
	}
	p := models.Product{ID: id}
	_,err = p.DeleteProduct(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			u.RespondWithError(w, http.StatusNotFound, "Product isn't found")
		default:
			u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	u.RespondWithJSON(w, http.StatusOK, "Element "+string(p.ID)+"deleted successful")

}

func (a *App) UpdateProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	uid, err := auth.ExtractTokenID(req)

	if err != nil {
		u.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	var p models.Product
	err = a.DB.QueryRow("select user_id from products where id=$1", id).Scan(&p.UserID)
	if err != nil {
		u.ERROR(w, http.StatusNotFound, errors.New("Product is not found"))
		return
	}

	if uid != p.UserID {
		u.ERROR(w, http.StatusUnauthorized, errors.New("It's notes of another user"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	if err := decoder.Decode(&p); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	p.UserID = uid
	p.Prepare()
	err = p.Validate()
	if err != nil {
		u.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	p.ID = id

	_,err = p.UpdateProduct(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			u.RespondWithError(w, http.StatusNotFound, "Product isn't found")
		default:
			u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	u.RespondWithJSON(w, http.StatusOK, "Element "+string(p.ID)+"updated successful")

}
