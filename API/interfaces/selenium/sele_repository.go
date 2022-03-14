package selenium

import (
	"strconv"
	"time"
)

type SeleniumRepository struct {
	SeleHandler
}

const (
	sign_in_URL = "https://sakito.cirkit.jp/user/sign_in"
)

type Data struct {
	Points      int
	BonusTicket int
	BonusWeek   string
}

// 所持ポイント数と所持ボーナス券枚数とボーナス券までの週を取得
func (repo *SeleniumRepository) Data() (*Data, error) {
	var data Data
	return &data, nil
}

func (repo *SeleniumRepository) BonusRoll(count int) error {
	return nil
}

func (repo *SeleniumRepository) CheckTicket() error {
	return nil
}

func (repo *SeleniumRepository) ExchangeRoll(count int) error {
	return nil
}

func findValueByElement(elem Element) (int, error) {
	var v int
	for {
		output, err := elem.Text()
		if err != nil {
			return 0, err
		}
		if output != "Waiting for remote server..." {
			// stringなのでintに変換
			v, err = strconv.Atoi(output)
			if err != nil {
				return 0, err
			}
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	return v, nil
}
