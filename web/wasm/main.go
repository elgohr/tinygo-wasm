package main

import (
	"syscall/js"
)

//go:generate tinygo build -o ../main.wasm -target wasm main.go
func main() {
	doc := js.Global().Get("document")
	doc.Call("getElementById", "some-field").Set("value", "initial value")
}

//go:export changeValue
func changeValue() {
	doc := js.Global().Get("document")
	doc.Call("getElementById", "some-field").Set("value", "changed value")
}
