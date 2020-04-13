package main_test

import (
	"github.com/siongui/godom"
	"testing"
)

func TestMain_SetsInitialValue(t *testing.T) {
	value := godom.Document.GetElementById("some-field").Value()
	if value != "initial value" {
		t.Errorf("should have set the field with a default value, but was %s", value)
	}
}
