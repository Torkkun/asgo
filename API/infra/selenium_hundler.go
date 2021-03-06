package infra

import (
	asgo "asgo/interfaces/selenium"

	"github.com/tebeka/selenium"
)

type SeleHandler struct {
	WebDriver selenium.WebDriver
}

type WebElement struct {
	Element selenium.WebElement
}

func NewSeleHandler() *SeleHandler {
	const (
		URL = "http://localhost:4444/wd/hub"
	)
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, URL)
	if err != nil {
		panic(err)
	}
	sele := new(SeleHandler)
	sele.WebDriver = wd
	return sele
}

func (driver *SeleHandler) Get(url string) error {
	err := driver.WebDriver.Get(url)
	return err
}

func (driver *SeleHandler) FindElement(by, value string) (asgo.Element, error) {
	elem, err := driver.WebDriver.FindElement(by, value)
	if err != nil {
		return new(WebElement), err
	}
	e := new(WebElement)
	e.Element = elem
	return e, nil
}

func (elem WebElement) Clear() error {
	return elem.Element.Clear()
}

func (elem WebElement) Click() error {
	return elem.Element.Click()
}

func (elem WebElement) SendKeys(keys string) error {
	return elem.Element.SendKeys(keys)
}

func (elem WebElement) Text() (string, error) {
	return elem.Element.Text()
}
