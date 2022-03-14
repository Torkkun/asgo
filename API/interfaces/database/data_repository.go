package database

import (
	"database/sql"
	"time"
)

type DataRepository struct {
	SqlHandler
}

type Data struct {
	UserID      string
	DailyPoint  int
	Points      int
	BonusTicket int
	BonusWeek   string
	UpdatedAt   time.Time
	// アンケートいるか分けるか
}

//データの管理
func (repo *DataRepository) InsertData(data *Data) error {
	_, err := repo.Execute(
		"INSERT INTO data(userid, daily_point, points, bonus_ticket, bonus_week) VALUES(?,?,?,?,?)",
		data.UserID, data.DailyPoint, data.Points, data.BonusTicket, data.BonusWeek)
	if err != nil {
		return err
	}
	return nil
}

func (repo *DataRepository) UpdateData(data *Data) error {
	_, err := repo.Execute(
		"UPDATE data SET daily_point = ?, points = ?, bonus_ticket = ?, bonus_week = ?",
		data.DailyPoint, data.Points, data.BonusTicket, data.BonusWeek)
	if err != nil {
		return err
	}
	return nil
}

func (repo *DataRepository) SelectDataByUserID(userID string) (*Data, error) {
	row := repo.QueryRow("SELECT * FROM data WHERE userid = ?")
	return convertToData(row)
}

func convertToData(row Row) (*Data, error) {
	data := Data{}
	err := row.Scan(&data.UserID, &data.DailyPoint, &data.Points, &data.BonusTicket, &data.BonusWeek, &data.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}
