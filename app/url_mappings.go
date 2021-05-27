package app

import (
	"github.com/go-sandbox/bookstore_users-api/controllers/ping"
	"github.com/go-sandbox/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// public
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/user/:user_id", users.Delete)

	// private
	router.GET("/internal/users/search", users.Search)
}
