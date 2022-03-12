package selenium

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

type SeleHandler interface {
	Get(string) error                            // Navigate to the URL.
	FindElement(string, string) (Element, error) // Get a reference to the text box.
}

type Element interface {
	Clear() error          // Clear clears the element.
	Click() error          // Click clicks on the element.
	SendKeys(string) error // SendKeys types into the element.
	Text() (string, error) // Text returns the text of the element.
}
