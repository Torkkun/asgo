package database

import (
	"asgo/domain"
	"time"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) InsertUser(u domain.UserDB) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO user(client_uid, firebase_uid, email, password, created_at, updated_at) values($1,$2,$3,$4) ON CONFLICT ON CONSTRAINT user_pkey UPDATE SET email = $3 & password = $4 & updated_at = current_timestamp",
		u.ClientUserID, u.FireBaseUserID, u.Email, u.Password,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) SelectUserFindById(identifier string) (user domain.UserDB, err error) {
	row, err := repo.Query(
		"SELECT client_uid, firebase_uid, email, password, created_at, updated_at FROM user WHERE client_uid = ?", identifier)
	if err != nil {
		return
	}
	defer row.Close()
	var client_uid string
	var firebase_uid string
	var email string
	var password string
	var created_at time.Time
	var updated_at time.Time
	row.Next()
	if err = row.Scan(&client_uid, &firebase_uid, &email, &password, &created_at, &updated_at); err != nil {
		return
	}
	user.ClientUserID = client_uid
	user.FireBaseUserID = firebase_uid
	user.Email = email
	user.Password = password
	user.CreatedAt = created_at
	user.UpdatedAt = updated_at
	return
}

func (repo *UserRepository) SelectUserFindByCode(identifier string) (secret domain.SecretCode, err error) {
	row, err := repo.Query("SELECT client_uid, onetimepass, created_at, updated_at FROM secret WHERE client_uid = ?", identifier)
	if err != nil {
		return
	}
	defer row.Close()
	var client_uid string
	var onetimepass string
	var created_at time.Time
	var updated_at time.Time
	if err = row.Scan(&client_uid, &onetimepass, &created_at, &updated_at); err != nil {
		return
	}
	secret.ClientUserID = client_uid
	secret.OneTimePassWord = onetimepass
	secret.CreatedAt = created_at
	secret.UpdatedAt = updated_at
	return
}
