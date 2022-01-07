package domain

//sakitoログイン用
type User struct {
	Email    string
	Password string
}

//ユーザー登録などの設定等の自前サイト用
type FireBaseUser struct {
}

type Users []User
