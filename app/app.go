package app

import (
	"fmt"

	"Rest/app/controllers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

var a = controllers.App{}

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	err=a.Initialize()
	check(err)
	a.Run(":8000")

}
func check(err error){
	if err!=nil {
		log.Fatal(err)
	}
}