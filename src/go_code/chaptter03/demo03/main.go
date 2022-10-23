package main

import (
	"fmt"
)

// 定义全局变量
var gn1 = 10
var gs1 = "zz"

// 第二种写法
var (
	gn2 = 20
	gs2 = "xx"
)

func main() {
	//多变量声明
	var n1, n2, n3 int
	fmt.Printf("n1: %v\n", n1)
	fmt.Printf("n2: %v\n", n2)
	fmt.Printf("n3: %v\n", n3)
	//多变量声明方式二  不建议这种方式 真的好不清晰
	var num, num1, name = 100, 200, "Sherry"
	fmt.Printf("num: %v\n", num)
	fmt.Printf("num1: %v\n", num1)
	fmt.Printf("name: %v\n", name)
	//类型推导方式
	b1, s1, i4 := true, "string", 10
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("i4: %v\n", i4)

	fmt.Printf("gn1: %v\n", gn1)
	fmt.Printf("gs1: %v\n", gs1)
	fmt.Printf("gn2: %v\n", gn2)
	fmt.Printf("gs2: %v\n", gs2)
}
