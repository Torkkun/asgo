package database

import (
	"asgo/domain"
	"crypto/sha256"
)

type SecretRepository struct {
	SqlHandler
}

func (repo *SecretRepository) InsertSecret(u domain.SecretCode) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO secret(client_uid, onetimepass) values($1,$2) ON CONFLICT ON CONSTRAINT user_pkey DO UPDATE SET onetimepass = $2 & updated_at = current_timestamp",
		u.ClientUserID, sha256.Sum256([]byte(u.OneTimePassWord)),
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
