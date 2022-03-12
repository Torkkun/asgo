package database

import (
	"time"
)

type UserRepository struct {
	SqlHandler
}

type User struct {
	Client_uid   string
	Firebase_uid string
	Email        string
	Password     string
	Created_at   time.Time
	Updated_at   time.Time
}

func (repo *UserRepository) InsertUser(user *User) error {
	_, err := repo.Execute(
		"INSERT INTO user(userid, firebase_uid, email, password, created_at, updated_at) values($1,$2,$3,$4) ON CONFLICT ON CONSTRAINT user_pkey UPDATE SET email = $3 & password = $4 & updated_at = current_timestamp",
		user.Client_uid, user.Firebase_uid, user.Email, user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) SelectUserFindById(identifier string) (*User, error) {
	row, err := repo.Query(
		"SELECT client_uid, firebase_uid, email, password, created_at, updated_at FROM user WHERE client_uid = ?", identifier)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var user User
	row.Next()
	if err = row.Scan(&user.Client_uid, &user.Firebase_uid, &user.Email, &user.Password, &user.Created_at, &user.Updated_at); err != nil {
		return nil, err
	}
	return &user, nil
}
