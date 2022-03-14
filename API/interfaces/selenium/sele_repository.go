package selenium

import (
	"time"
)

type SeleniumRepository struct {
	SeleHandler
}

const (
	sign_in_URL = "https://sakito.cirkit.jp/user/sign_in"
)

func (repo *SeleniumRepository) BonusRoll(count int) error {
	return nil
}

func (repo *SeleniumRepository) CheckTicket() error {
	return nil
}

func (repo *SeleniumRepository) ExchangeRoll(count int) error {
	return nil
}

// テキストを取得する
func findValueByElement(elem Element) (string, error) {
	var v string
	for {
		output, err := elem.Text()
		if err != nil {
			return "", err
		}
		if output != "Waiting for remote server..." {
			v = output
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	return v, nil
}
