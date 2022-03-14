package database

import "database/sql"

type ClientRepository struct {
	SqlHandler
}

// ユーザーのclientidを格納
type Client struct {
	UserID       string //firebaseのID
	ClientUserID string
}

func (repo *ClientRepository) InsertClient(client *Client) error {
	_, err := repo.Execute("INSERT INTO client(userid, client_uid) VALUES(?,?)",
		client.UserID, client.ClientUserID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ClientRepository) SelectClientFindByID(userID string) (*Client, error) {
	row := repo.QueryRow("")
	return convertToClient(row)
}

func convertToClient(row Row) (*Client, error) {
	client := Client{}
	err := row.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &client, nil
}
