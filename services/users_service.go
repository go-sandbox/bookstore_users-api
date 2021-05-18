package services

import (
	"net/http"

	"github.com/go-sandbox/bookstore_users-api/domain/users"
	"github.com/go-sandbox/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
