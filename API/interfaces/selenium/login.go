package selenium

//sakitoログイン用
type Login struct {
	Email    string
	Password string
}

// 自動ログイン処理を実行
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
	// クリックして終了
	return elem.Click()

}
