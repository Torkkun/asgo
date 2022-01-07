package selenium

import (
	"github.com/tebeka/selenium"
)

func SettingSelenium() {
	const (
		URL = "http://localhost:4444/wd/hub"
	)
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, URL)
	if err != nil {
		panic(err)
	}
	defer wd.Quit()
}
