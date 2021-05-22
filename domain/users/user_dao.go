package users

import (
	"fmt"
	"strings"

	"github.com/go-sandbox/bookstore_users-api/datasources/mysql/users_db"
	"github.com/go-sandbox/bookstore_users-api/utils/date_utils"
	"github.com/go-sandbox/bookstore_users-api/utils/errors"
)

const (
	// TODO: MySQL固有のメッセージなのでinfra層に移動する
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"

	// SQL
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_Name, email, date_created FROM users WHERE id = ?;"
)

func (user User) Get() *errors.RestErr {
	// クエリ読み込み
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// 検索処理
	result := stmt.QueryRow(user.Id)
	// 取得できなかった場合
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		// 取得できなかった理由：存在しなかった
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user User) Save() *errors.RestErr {

	// クエリ読み込み
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// 現在時刻取得
	user.DateCreated = date_utils.GetNowString()
	// 登録処理
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		// メールアドレスのユニークキー違反エラーの場合
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError("email %s already exist")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	// ユーザーID確認
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId

	return nil
}
