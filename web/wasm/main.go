package main

import (
	"github.com/siongui/godom"
)

//go:generate tinygo build -o ../main.wasm -target wasm main.go
//go:generate sh -ec "go build ../../vendor/github.com/agnivade/wasmbrowsertest/"
//go:generate mv wasmbrowsertest $GOPATH/bin/go_js_wasm_exec"
func main() {
	godom.Document.GetElementById("some-field").SetValue("initial value")
}
