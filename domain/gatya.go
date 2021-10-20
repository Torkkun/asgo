package domain

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

type User struct {
	Email    string
	Password string
}

type Gatyas []Gatya

type Users []User
