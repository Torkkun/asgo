package database

import (
	"crypto/sha256"
	"database/sql"
	"time"
)

type SecretRepository struct {
	SqlHandler
}

type Secret struct {
	ClientUserID    string //クライアントごとのID
	OneTimePassWord string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (repo *SecretRepository) InsertSecret(user *Secret) error {
	_, err := repo.Execute(
		"INSERT INTO secret(client_uid, onetimepass) values($1,$2)",
		user.ClientUserID, sha256.Sum256([]byte(user.OneTimePassWord)),
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SecretRepository) SelectSecretById(userID string) (*Secret, error) {
	row := repo.QueryRow("")
	return convertToSecret(row)
}

func convertToSecret(row Row) (*Secret, error) {
	secret := Secret{}
	err := row.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &secret, nil
}
