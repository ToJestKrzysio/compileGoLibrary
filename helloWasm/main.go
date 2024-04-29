package main

import (
    "fmt"
    "syscall/js"
)


func hello(this js.Value, args []js.Value) any {
    fmt.Println("Hello world!")
    return nil
}

func main() {
    c := make(chan struct{}, 0)
    js.Global().Set("hello", js.FuncOf(hello))
    <-c

    alert := js.Global().Get("alert")
	alert.Invoke("Hi")
}
