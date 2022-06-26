package main

import (
	"fmt"
	"sync"
	"syscall/js"

	"github.com/jiro4989/ojosama"
)

var wg sync.WaitGroup

func convert(this js.Value, args []js.Value) interface{} {
	text, err := ojosama.Convert(args[0].String(), nil)
	if err != nil {
		return ""
	}
	return text
}

func registerCallbacks() {
	defer wg.Done()
	chjs := js.FuncOf(convert)
	js.Global().Set("convert", chjs)
}

func main() {
	wg.Add(1)
	println("[LOG] WebAssembly Initialized")
	go registerCallbacks()
	wg.Wait()
	fmt.Println("[LOG] Converting")
}
