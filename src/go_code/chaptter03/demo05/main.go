package main

import "fmt"

//程序中+号的使用
func main() {

	var i = 1
	var j = 2
	r := i + j //运算
	fmt.Printf("r: %v\n", r)

	s1 := "hello"
	s2 := "world"
	s3 := s1 + s2 //拼接
	fmt.Printf("s3: %v\n", s3)
}
