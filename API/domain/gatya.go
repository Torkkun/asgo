package domain

import "time"

type Ticket struct {
	Name    string
	HoldNum int
	Period  string
}

type Gatya struct {
	Point          int
	Tickets        Ticket
	BonusTicket    int
	BonusCountTime string
}

type DailyGatya struct {
	Point          int
	Point_Sum      int
	Roll_num       int
	Execution_Date time.Time
}

type Gatyas []Gatya
