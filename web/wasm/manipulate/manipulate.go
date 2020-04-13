package manipulate

import "github.com/siongui/godom"

//go:export ChangeValue
func ChangeValue() {
	godom.Document.GetElementById("some-field").SetValue("changed value")
}