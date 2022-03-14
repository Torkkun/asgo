package selenium

import "time"

type DailyGatya struct {
	DayPoint      int
	PointSum      int
	ExecutionDate string
}

// デイリーを回して情報をスクレイピング
func (repo *SeleniumRepository) DailyRoll() (*DailyGatya, error) {
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

	// 少し待つ
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
	date := time.Now().Format("2006/01/02 15:04:05")
	// 最後に入れる
	return &DailyGatya{
			DayPoint:      point,
			PointSum:      sum,
			ExecutionDate: date},
		nil
}
