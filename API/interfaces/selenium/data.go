package selenium

import "fmt"

type Data struct {
	Points      int
	BonusTicket int
	BonusWeek   string
}

// 所持ポイント数と所持ボーナス券枚数とボーナス券までの週を取得
func (repo *SeleniumRepository) Data() (*Data, error) {
	// 結果がどうなるかまた狙った値が取得されない場合のエラーハンドリングをどうするか
	// 所持ポイント
	elem, err := repo.FindElement(ByCSSSelector, "body > div > div:nth-child(1) > div:nth-child(1) > div > div > div > div.col-xs-9.text-right")
	if err != nil {
		return nil, err
	}
	point, err := findValueByElement(elem)
	if err != nil {
		return nil, err
	}
	fmt.Println(point)

	// 所持ボーナス券
	elem, err = repo.FindElement(ByCSSSelector, "body > div > div:nth-child(1) > div:nth-child(2) > div > div > div > div.col-xs-9.text-right")
	if err != nil {
		return nil, err
	}
	bonusTicket, err := findValueByElement(elem)
	if err != nil {
		return nil, err
	}
	fmt.Println(bonusTicket)

	// ボーナス券まで
	elem, err = repo.FindElement(ByCSSSelector, "body > div > div:nth-child(1) > div:nth-child(3) > div > div > div > div:nth-child(3) > h1")
	if err != nil {
		return nil, err
	}
	bonusWeek, err := findValueByElement(elem)
	if err != nil {
		return nil, err
	}
	fmt.Println(bonusWeek)
	return nil, nil
}
