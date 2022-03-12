package database

import (
	"crypto/sha256"
	"time"
)

type SecretRepository struct {
	SqlHandler
}

type Secret struct {
	ClientUserID    string
	OneTimePassWord string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (repo *SecretRepository) InsertSecret(user *Secret) error {
	_, err := repo.Execute(
		"INSERT INTO secret(client_uid, onetimepass) values($1,$2) ON CONFLICT ON CONSTRAINT user_pkey DO UPDATE SET onetimepass = $2 & updated_at = current_timestamp",
		user.ClientUserID, sha256.Sum256([]byte(user.OneTimePassWord)),
	)
	if err != nil {
		return err
	}
	return nil
}
