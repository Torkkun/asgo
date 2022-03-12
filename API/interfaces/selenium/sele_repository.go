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

//sakitoログイン用
type Login struct {
	Email    string
	Password string
}

type DailyGatya struct {
	Day_Point      int
	Point_Sum      int
	Execution_Date string
}

//自動ログイン処理を実行
func (repo *SeleniumRepository) Login(login *Login) (err error) {
	if err = repo.Get(sign_in_URL); err != nil {
		return
	}
	elem, err := repo.FindElement(ByCSSSelector, "#user_email")
	if err != nil {
		return
	}
	if err = elem.SendKeys(login.Email); err != nil {
		return
	}
	elem, err = repo.FindElement(ByCSSSelector, "#user_password")
	if err != nil {
		return
	}
	if err = elem.SendKeys(login.Password); err != nil {
		return
	}
	elem, err = repo.FindElement(ByCSSSelector, ".btn.btn-info.btn-block")
	if err != nil {
		return
	}
	if err = elem.Click(); err != nil {
		return
	}
	elem, err = repo.FindElement(ByCSSSelector, ".btn-primary")
	if err != nil {
		return
	}
	//クリックして終了
	return elem.Click()

}

//デイリーを回して情報をスクレイピング
func (repo *SeleniumRepository) DailyRoll() (*DailyGatya, error) {
	date := time.Now().Format("2006/01/02 15:04:05")
	elem, err := repo.FindElement(ByLinkText, "ポイントガチャを回す")
	if err != nil {
		return nil, err
	}
	if err = elem.Click(); err != nil {
		return nil, err
	}
	elem, err = repo.FindElement(ByCSSSelector, ".btn.btn-success.btn-block")
	if err != nil {
		return nil, err
	}
	if err = elem.Click(); err != nil {
		return nil, err
	}

	//少し待つ
	time.Sleep(time.Millisecond * 100)

	elem, err = repo.FindElement(ByTagName, "b")
	if err != nil {
		return nil, err
	}
	point, err := findValueByElement(elem)
	if err != nil {
		return nil, err
	}
	elem, err = repo.FindElement(ByTagName, "h3")
	if err != nil {
		return nil, err
	}
	sum, err := findValueByElement(elem)
	if err != nil {
		return nil, err
	}
	elem, err = repo.FindElement(ByLinkText, "引換券ガチャへ")
	if err != nil {
		return nil, err
	}
	if err = elem.Click(); err != nil {
		return nil, err
	}
	//最後に入れる
	return &DailyGatya{Day_Point: point, Point_Sum: sum, Execution_Date: date}, nil
}

func (repo *SeleniumRepository) BonusRoll(count int) (err error) {
	return
}

func (repo *SeleniumRepository) CheckTicket() (err error) {
	return
}

func (repo *SeleniumRepository) ExchangeRoll(count int) (err error) {
	return
}

func findValueByElement(elem Element) (int, error) {
	var v int
	for {
		output, err := elem.Text()
		if err != nil {
			return 0, err
		}
		if output != "Waiting for remote server..." {
			//stringなのでintに変換
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
