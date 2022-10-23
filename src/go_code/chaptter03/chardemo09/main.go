package main

import "fmt"

func main() {
	//go的字符串是由单个字节链接起来的
	//传统的字符串是由字符组成的 而go的字符串不同它是由字节组成的
	//golang中没有专门的字符类型 如果要存储单个字符 一般使用byte来保存
	//官方将string归属到基本数据类型

	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %v\n", c2)
}
