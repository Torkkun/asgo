package domain

import "time"

//引換券のデータ
type ExchangeList struct {
	Tickets []Ticket
}

type Ticket struct {
	Name    string
	HoldNum int
	Period  string
}

//データベース内のガチャ結果
type Gatya struct {
	Day_Point       int
	Sum_Point       int
	Bonus_Ticket    int
	Until_Bonus     int
	Exchange_Ticket int
	Enquete         bool
	User_ID         string
	Updated_At      time.Time
}

//デイリー用のデータ
type DailyGatya struct {
	Day_Point      int
	Point_Sum      int
	Execution_Date time.Time
}

//引換券ガチャ用のデータ
type ExchangeGatya struct {
}

//ボーナスガチャのデータ
type BonusGatya struct {
}
