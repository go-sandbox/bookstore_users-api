package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sandbox/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run("127.0.0.1:8080")
}
