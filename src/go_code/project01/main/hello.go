package main

import "fmt"

const (
	c1 = 2
	c2
	c3
	c4 = 6
	c5
)

func main() {

	var a, b int = 10, 5
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	a, b = b, a // 交换ab值
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %v\n", c2)
	fmt.Printf("c3: %v\n", c3)
	fmt.Printf("c4: %v\n", c4)
	fmt.Printf("c5: %v\n", c5)
	//1.go扩展名
	//2.入口为main方法
	//3.严格区分大小写
	//4.不需要分号
	//5.定义的变量或者 包没有用到 编译不通过
	//6.大括号成对出现
	//7.go build -o 文件名.exe 文件.go 指定可执行文件名
	//8.转义字符: \t 制表符tab \n换行 \\ 一个\   \" 一个"  \r 回车
	fmt.Print("hello world\n")

	str := "hello, world"
	str1 := str[:5]  // 获取索引5（不含）之前的子串
	str2 := str[7:]  // 获取索引7（含）之后的子串
	str3 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("str3:", str3)
}
