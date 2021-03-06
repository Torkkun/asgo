package domain

//引換券のデータ
type ExchangeList struct {
	Tickets []Ticket
}

type Ticket struct {
	Name    string
	HoldNum int
	Period  string
}

//デイリー用のデータ
type DailyGatyaResponse struct {
	DayPoint      int    `json:"day_point"`
	PointSum      int    `json:"point_sum"`
	ExecutionDate string `json:"execution_data"`
}

//引換券ガチャ用のデータ
type ExchangeGatya struct {
}

//ボーナスガチャのデータ
type BonusGatya struct {
}

// データの取得
type DataResponse struct {
	DayPoint    int    `json:"day_point"`
	Points      int    `json:"points"`
	BonusTicket int    `json:"bonus_ticket"`
	BonusWeek   string `json:"bonus_week"`
}
