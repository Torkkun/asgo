package infra

import (
	"log"
	"os"

	"github.com/tebeka/selenium"
)

var (
	appURL = "http://play.golang.org/?simple=1"
	//wd     selenium.WebDriver
)

func Letstartsele() {

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "")

	if err != nil {
		log.Println("Failed to start selenium:", err.Error())
		os.Exit(1)
	}
	defer wd.Quit()
	wd.Get(appURL)

}
