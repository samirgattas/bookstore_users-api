package app

import (
	"github.com/develop-microservices-in-go/bookstore_users-api/controllers/ping"
	"github.com/develop-microservices-in-go/bookstore_users-api/controllers/users"
)

func mapUrls() {
	// Endpoint used by many servers (AWS, GCP, etc) to check if API is alive
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
