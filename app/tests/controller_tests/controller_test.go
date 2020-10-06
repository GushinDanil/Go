package controller_tests

import (
	"Rest/app/controllers"
	"Rest/app/models"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)
func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}
var appTest = controllers.App{}

func TestMain(m *testing.M){


	Database()

	os.Exit(m.Run())
}
func Database(){
	var err error
	connectionString:=  fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))

	appTest.DB ,err = sql.Open("postgres",connectionString)
	if err!=nil {
		fmt.Printf("Cannot connect to  database\n")
		log.Fatal("This is the error:", err)
	}

}

func seedOneUser() (models.User, error) {

	_,err := appTest.DB.Exec("truncate table users")
	if err != nil {
		log.Fatal(err)
	}

	u := models.User{
		Nickname: "Pet",
		Email:    "pet@gmail.com",
		Password: "password",
	}
	err = appTest.DB.QueryRow("insert into  users (nickname,email,password)values ($1,$2,$3) Returning id", u.Nickname, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func seedUsers() ([]models.User , error) {

	u := []models.User{
		models.User{
			Nickname: "Steven victor",
			Email:    "steven@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Kenny Morris",
			Email:    "kenny@gmail.com",
			Password: "password",
		},
	}
	for i:= range u {
		err := appTest.DB.QueryRow("insert into  users (nickname,email,password)values ($1,$2,$3) Returning id", u[i].Nickname, u[i].Email, u[i].Password).Scan(&u[i].ID)


		if err != nil {
			return  nil,err
		}
	}
	return  u,nil
}
