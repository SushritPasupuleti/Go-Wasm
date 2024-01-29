package main

import (
	"fmt"
	"syscall/js" // This import may not be available in your editor, 
	// but it will compile on >1.11
)

// func add(a, b int) int {
//
// 	return a + b
// }

func add(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()

	return a + b
}

func main() {
	fmt.Println("Hello from Go!")

	js.Global().Set("add", js.FuncOf(add)) // Expose the add function to JavaScript

	<-make(chan bool) // Wait forever to avoid exiting the program prematurely
}
