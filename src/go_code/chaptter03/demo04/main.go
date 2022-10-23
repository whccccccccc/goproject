package main

import "fmt"

//变量使用的注意事项
func main() {
	//该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	// i = 1.1//不对的 不能改变类型
	//变量在同一个作用域内不能重名
	// var i int=20//不对的
	//变量=变量名+值+数据类型  变量三要素
	//没有赋初始值 都会有默认值
	fmt.Printf("i: %v\n", i)
}
