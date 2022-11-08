package main

import (
	"fmt"
)

func main() {
	var ptr *int
	fmt.Printf("ptr: %v\n", ptr)
	a := 32
	ptr = &a
	fmt.Printf("ptr: %v\n", ptr)
	fmt.Printf("ptr: %v\n", *ptr)
}
