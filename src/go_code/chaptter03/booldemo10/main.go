package main

import (
	"fmt"
	"unsafe"
)

// bool类型基本介绍
// 1. 只能取 true和false
// 2. 占用1字节
func main() {

	var b bool = false
	fmt.Printf("b: %v\n", b)
	//占用空间
	fmt.Printf("占用空间: %v\n", unsafe.Sizeof(b))
	//不可以用0或者非0的整数来代替
}
