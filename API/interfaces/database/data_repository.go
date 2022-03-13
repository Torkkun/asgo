package database

type DataRepository struct {
	SqlHandler
}

type Data struct {
	UserID      string
	DailyPoint  int
	Points      int
	BonusTicket int
	BonusWeek   string
	// アンケートいるか分けるか
}

//ユーザーデータの管理はfirebase自分で持つのはユーザーごとのデータ
func (repo *DataRepository) InsertData() {
}
