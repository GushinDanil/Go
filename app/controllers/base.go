package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() error {

	var err error
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	a.DB, err = sql.Open("postgres", connectionString)
	check(err)

	a.Router = mux.NewRouter()

	a.InitializeRoutes()
	fmt.Println("Connected")

	return nil
}
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
