package database

import (
	"database/sql"
	"time"
)

type UserRepository struct {
	SqlHandler
}

type User struct {
	UserID     string //firebaseのID
	Email      string //sakitoのメールアドレス
	Password   string //sakitoのパスワード
	Created_at time.Time
	Updated_at time.Time
}

func (repo *UserRepository) InsertUser(user *User) error {
	_, err := repo.Execute(
		"INSERT INTO user(userid, email, password) values(?,?,?)",
		user.UserID, user.Email, user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) SelectUserFindById(userID string) (*User, error) {
	row := repo.QueryRow(
		"SELECT userid, email, password, created_at, updated_at FROM user WHERE userid = ?", userID)
	return convertToUser(row)
}

func convertToUser(row Row) (*User, error) {
	user := User{}
	err := row.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
