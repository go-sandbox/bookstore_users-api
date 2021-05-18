package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sandbox/bookstore_users-api/domain/users"
	"github.com/go-sandbox/bookstore_users-api/services"
	"github.com/go-sandbox/bookstore_users-api/utils/errors"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User

	// フォーマットエラー
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	// 登録エラー
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
