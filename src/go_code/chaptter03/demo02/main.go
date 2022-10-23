package main

import "fmt"

func main() {
	//1.指定变量类型声明后不赋值  使用默认值
	var i int
	fmt.Printf("i: %v\n", i)
	//2.根据值自行判断变量类型 类型推导
	var num = 10
	fmt.Printf("num: %v\n", num)
	//3. 省略var  :=左侧不可以是已经声明过的变量
	name := "Sherry"
	fmt.Printf("name: %v\n", name)

}
