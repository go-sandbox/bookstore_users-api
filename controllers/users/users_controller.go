package users

import (
	"net/http"
	"strconv"

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
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	// 取得エラー
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	// ユーザー確認
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	// フォーマットエラー
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	// PATCHの場合
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.CreateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
