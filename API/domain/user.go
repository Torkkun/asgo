package domain

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
	Email     string
	DiscordID string
	Password  string
}

type UsersDB []UserDB
