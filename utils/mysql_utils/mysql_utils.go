package mysql_utils

import (
	"strings"

	"github.com/go-sandbox/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	// クエリエラーの場合
	if !ok {
		// 検索結果がない場合
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	// 参考：https://dev.mysql.com/doc/refman/5.6/ja/error-messages-server.html
	case 1062:
		// ユニークキー違反の場合
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
