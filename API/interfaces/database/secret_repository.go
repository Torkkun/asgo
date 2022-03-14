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

func (repo *SecretRepository) InsertSecret(secret *Secret) error {
	_, err := repo.Execute(
		"INSERT INTO secret(client_uid, onetimepass) values(?,?)",
		secret.ClientUserID, sha256.Sum256([]byte(secret.OneTimePassWord)),
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SecretRepository) SelectSecretById(clientUserID string) (*Secret, error) {
	row := repo.QueryRow("SELECT * FROM secret WHERE client_uid = ?", clientUserID)
	return convertToSecret(row)
}

func convertToSecret(row Row) (*Secret, error) {
	secret := Secret{}
	err := row.Scan(&secret.ClientUserID, &secret.OneTimePassWord, &secret.CreatedAt, &secret.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &secret, nil
}
