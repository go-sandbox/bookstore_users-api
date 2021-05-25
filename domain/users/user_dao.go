package users

import (
	"github.com/go-sandbox/bookstore_users-api/datasources/mysql/users_db"
	"github.com/go-sandbox/bookstore_users-api/utils/date_utils"
	"github.com/go-sandbox/bookstore_users-api/utils/errors"
	"github.com/go-sandbox/bookstore_users-api/utils/mysql_utils"
)

const (
	// SQL
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_Name, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id = ?;"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(getErr)
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
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	// ユーザーID確認
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}

	user.Id = userId

	return nil
}

func (user User) Update() *errors.RestErr {
	// クエリ読み込み
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// 更新処理
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user User) Delete() *errors.RestErr {
	// クエリ読み込み
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// 削除処理
	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}
