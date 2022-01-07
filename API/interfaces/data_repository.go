package interfaces

type DataRepository struct {
	SqlHandler
}

//ユーザーデータの管理はfirebase自分で持つのはユーザーごとのデータ
func (repo *DataRepository) Store() {
}
