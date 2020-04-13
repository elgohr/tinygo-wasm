package manipulate_test

import (
	"github.com/elgohr/tinygo-wasm/web/wasm/manipulate"
	"github.com/siongui/godom"
	"testing"
)

func TestChangeValueChangesValue(t *testing.T) {
	initial := godom.Document.GetElementById("some-field").Value()
	if initial != "initial value" {
		t.Errorf("should have set the field with a default value, but was %s", initial)
	}
	manipulate.ChangeValue()
	changed := godom.Document.GetElementById("some-field").Value()
	if changed != "changed value" {
		t.Errorf("should have set the field with a changed value, but was %s", changed)
	}
}
