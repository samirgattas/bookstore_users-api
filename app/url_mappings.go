package app

import (
	"github.com/develop-microservices-in-go/bookstore_users-api/controllers/ping"
	"github.com/develop-microservices-in-go/bookstore_users-api/controllers/users"
)

func mapUrls() {
	// Endpoint used by many servers (AWS, GCP, etc) to check if API is alive
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.GET("/users/search", users.Search)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
