package database

import "asgo/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.UserDB) (id int, err error) {
	result, err := repo.Execute(
		"", u.Email, u.Password, u.DiscordID,
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

func (repo *UserRepository) FindById(identifier int) (user domain.UserDB, err error) {
	row, err := repo.Query("")
	if err != nil {
		return
	}
	defer row.Close()
	var email string
	var password string
	var discord_id string
	row.Next()
	if err = row.Scan(&email, &password, &discord_id); err != nil {
		return
	}
	user.Email = email
	user.Password = password
	user.DiscordID = discord_id
	return
}

func (repo *UserRepository) FindAll() (users domain.UsersDB, err error) {
	rows, err := repo.Query("")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var email string
		var password string
		var discord_id string
		if err := rows.Scan(&email, &password, &discord_id); err != nil {
			continue
		}
		user := domain.UserDB{
			Email:     email,
			Password:  password,
			DiscordID: discord_id,
		}
		users = append(users, user)
	}
	return
}
