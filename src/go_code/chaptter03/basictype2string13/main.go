package main

import (
	"fmt"
	"strconv"
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

	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str: %v type:%T\n", str, str)
	//f代表格式 10代表精度控制 10位小数位 不足0补齐 64代表这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str: %v type:%T\n", str, str)
	str = strconv.FormatBool(b2)
	fmt.Printf("str: %v type:%T\n", str, str)

	str2 := strconv.Itoa(32)
	fmt.Printf("str2: %v type: %T \n", str2, str2)
}
