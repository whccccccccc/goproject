package main

import (
	"fmt"
)

/*
基本介绍:
1. 基本数据类型变量存的就是值 也叫值类型
2. 获取变量的地址用 &变量名 这种写法
3. 指针类型,指针变量存的是一个地址,这个地址指向的空间存的才是值
4. 获取指针类型所指向的值, 使用 *指针变量名
*/
func main() {
	var i int = 10
	//i的地址
	fmt.Println("i的地址:", &i) //0xc000016098
	//ptr 是一个指针变量
	//ptr的类型 *int
	//ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr: %v\n", ptr)
	fmt.Printf("ptr指向的值: %v\n", *ptr)
	fmt.Printf("i: %v\n", i)
	*ptr = 11
	fmt.Printf("i: %v\n", i)
	//注:指针也区分类型  *类型A  要和 类型B&  对应
	//值类型都有对应的指针类型
	//值类型包括 基本数据类型 数组 和结构体
	//引用类型 指针 slice 切片 map  管道 interface
}
