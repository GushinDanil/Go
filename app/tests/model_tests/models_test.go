package model_tests
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
"time"
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

func seedUsers() error {

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
	for i := range u {
		err := appTest.DB.QueryRow("insert into  users (nickname,email,password)values ($1,$2,$3) Returning id", u[i].Nickname, u[i].Email, u[i].Password).Scan(&u[i].ID)


		if err != nil {
			return  err
		}
	}
	return  nil
}

func seedOneUserAndOneProduct() (models.Product, error) {

	_,err := appTest.DB.Exec("truncate table users")
	if err != nil {
		log.Fatal(err)
	}
	_,err = appTest.DB.Exec("truncate table products")
	if err != nil {
		log.Fatal(err)
	}
	u := models.User{
		Nickname: "Sam Phil",
		Email:    "sam@gmail.com",
		Password: "password",
	}
	err = appTest.DB.QueryRow("insert into  users (nickname,email,password)values ($1,$2,$3) Returning id", u.Nickname, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return models.Product{}, err
	}
	p := models.Product{
		Name:    "Pumpkin",
		Price:  30,
		UserID: int(u.ID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = appTest.DB.QueryRow("insert into  products (name,price,user_id,createdat,updatedat)values ($1,$2,$3,$4,$5) Returning id", p.Name, p.Price,p.UserID,p.CreatedAt,p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		return models.Product{}, err
	}
	return p, nil
}

func seedUsersAndProducts() ([]models.User, []models.Product, error) {


	var u = []models.User{
		models.User{
			Nickname: "Gushchin Danil",
			Email:    "danilgusin@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Ivanov Viktor",
			Email:    "nikita@gmail.com",
			Password: "password",
		},
	}
	var p = []models.Product{
		models.Product{
			Name:   "orange",
			Price: 23,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.Product{
			Name:   "tomato",
			Price: 20,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for i  := range u {
		err := appTest.DB.QueryRow("insert into  users (nickname,email,password)values ($1,$2,$3) Returning id", u[i].Nickname, u[i].Email, u[i].Password).Scan(&u[i].ID)
		if err != nil {
			log.Fatalf("cannot seed products table: %v", err)
		}
		p[i].UserID=int(u[i].ID)
		err = appTest.DB.QueryRow("insert into  products (name,price,user_id,createdat,updatedat)values ($1,$2,$3,$4,$5) Returning id", p[i].Name, p[i].Price,p[i].UserID,p[i].CreatedAt,p[i].UpdatedAt).Scan(&p[i].ID)
		if err!=nil {
			return nil,nil,err
		}

	}
	return u, p, nil
}
