package domain

import "time"

//sakitoログイン用
type User struct {
	Email    string
	Password string
}

type UserResponce struct {
}

type UsersRequest struct {
	Users []User `json:"users"`
}

//ユーザー登録などの設定等の自前サイト用
type FireBaseUser struct {
}

type UserDB struct {
	ClientUserID   string
	FireBaseUserID string
	Email          string
	Password       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SecretCode struct {
	ClientUserID    string
	OneTimePassWord string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Code struct {
	Code string `json:"code"`
}
