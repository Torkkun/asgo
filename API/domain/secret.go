package domain

import "time"

type SecretCode struct {
	ClientUserID    string
	OneTimePassWord string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
