package main

import (
	"fmt"
)

// this is the only way to initialize outside a func
var (
	name string = "santiago"
	py   bool
)

func main() {
	fmt.Printf("Type: %t String: %q", py, name)
}
