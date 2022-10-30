package main

import (
	"fmt"
	_ "unsafe"
)

// 基本数据类型 转换string
func main() {

	var n1 int = 99

	var n2 float64 = 23.5656
	// var b bool = false
	// var myChar byte = 'h'
	var str string
	//第一种方法  fmt.Sprintf方法
	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str: %q\n", str)
	fmt.Printf("str: %T\n", str)

	str = fmt.Sprintf("%f", n2)
	fmt.Printf("str: %q\n", str)
	fmt.Printf("str: %T\n", str)
	//第二种方法使用strcorv包的函数
}
