package selenium

import (
	"asgo/domain"
	"strconv"
	"time"
)

type SeleniumRepository struct {
	SeleHandler
}

const (
	Sign_in_URL = "https://sakito.cirkit.jp/user/sign_in"
)

const (
	ByID              = "id"
	ByXPATH           = "xpath"
	ByLinkText        = "link text"
	ByPartialLinkText = "partial link text"
	ByName            = "name"
	ByTagName         = "tag name"
	ByClassName       = "class name"
	ByCSSSelector     = "css selector"
)

//自動ログイン処理を実行
func (repo *SeleniumRepository) Login(user domain.User) (err error) {
	if err = repo.Get(Sign_in_URL); err != nil {
		return
	}
	elem, err := repo.FindElement(ByCSSSelector, "#user_email")
	if err != nil {
		return
	}
	if err = elem.SendKeys(user.Email); err != nil {
		return
	}
	elem, err = repo.FindElement(ByCSSSelector, "#user_password")
	if err != nil {
		return
	}
	if err = elem.SendKeys(user.Password); err != nil {
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
func (repo *SeleniumRepository) DailyRoll() (daily domain.DailyGatya, err error) {
	var point int
	var sum int
	var output string

	date := time.Now().Format("2006/01/02 15:04:05")

	elem, err := repo.FindElement(ByLinkText, "ポイントガチャを回す")
	if err != nil {
		return
	}
	if err = elem.Click(); err != nil {
		return
	}
	elem, err = repo.FindElement(ByCSSSelector, ".btn.btn-success.btn-block")
	if err != nil {
		return
	}
	if err = elem.Click(); err != nil {
		return
	}

	//少し待つ
	time.Sleep(time.Millisecond * 100)

	elem, err = repo.FindElement(ByTagName, "b")
	if err != nil {
		return
	}
	for {
		output, err = elem.Text()
		if err != nil {
			return
		}
		if output != "Waiting for remote server..." {
			//stringなのでintに変換
			point, err = strconv.Atoi(output)
			if err != nil {
				return
			}
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	elem, err = repo.FindElement(ByTagName, "h3")
	if err != nil {
		return
	}
	for {
		output, err = elem.Text()
		if err != nil {
			return
		}
		if output != "Waiting for remote server..." {
			//stringなのでintに変換
			sum, err = strconv.Atoi(output)
			if err != nil {
				return
			}
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	elem, err = repo.FindElement(ByLinkText, "引換券ガチャへ")
	if err != nil {
		return
	}
	if err = elem.Click(); err != nil {
		return
	}
	//最後に入れる

	daily.Day_Point = point
	daily.Point_Sum = sum
	daily.Execution_Date = date
	return
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
