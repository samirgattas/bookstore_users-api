package main

import (
	"log"

	"github.com/develop-microservices-in-go/bookstore_users-api/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	app.StartApplication()
}
