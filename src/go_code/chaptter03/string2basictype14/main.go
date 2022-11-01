package main

import (
	"fmt"
	"strconv"
)

// string 转换为基本数据类型
func main() {

	var strB string = "true"

	//排除函数的第二个返回值
	b, _ := strconv.ParseBool(strB)
	fmt.Printf("b: %v type:%T \n", b, b)

	var strI string = "300"
	//字符串,进制,长度  int64
	i, _ := strconv.ParseInt(strI, 10, 64)
	fmt.Printf("i: %v type:%T \n", i, i)

	var strF string = "30.236"
	//字符串,精度 32/64
	f, _ := strconv.ParseFloat(strF, 64)
	fmt.Printf("f: %v type:%T \n", f, f)

	//注意 异常数据转换时 会直接转换成 默认值  例如 hello 转换int 结果为0

	var hello string = "hello"

	n, _ := strconv.ParseInt(hello, 10, 64)
	fmt.Printf("n: %v\n", n)

}
