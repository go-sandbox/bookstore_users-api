package users

import (
	"fmt"
	"strings"

	"github.com/go-sandbox/bookstore_users-api/datasources/mysql/users_db"
	"github.com/go-sandbox/bookstore_users-api/utils/date_utils"
	"github.com/go-sandbox/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"

	// SQL
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	// モック
	usersDB = make(map[int64]*User)
)

func something() {
	user := User{}
	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.FirstName)
}

func (user User) Get() *errors.RestErr {
	// DB疎通確認
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d not found.", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

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
