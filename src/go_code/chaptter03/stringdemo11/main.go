package main

import "fmt"

func main() {
	//1.Go的字符串是由单个字节连接起来的Go语言的字符串的字节使用UTF-8编码表示Unicode文本
	var address string = "郑州"
	fmt.Printf("address: %v\n", address)
	//使用细节

	//2.字符串一旦赋值了  字符串不能修改 字符串不可变性  和Java一样
	var s1 string = "hello"
	// s1[0] = "zzz"
	s1 = "zzzzz" //但是这里又可以修改.....
	fmt.Printf("s1: %v\n", s1)
	//3.两种表示形式
	//3.1 双引号  会识别转义字符
	//3.2 反引号  以字符的原生形式输出 包括换行和特殊字符 可以实现防止攻击

	s2 := "zzz\nzz"
	s3 := `zzz\nzz`
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	//4.字符串的拼接方式
	var s4 = "hello" + "world"
	fmt.Printf("s4: %v\n", s4)
	s4 += "haha"
	fmt.Printf("s4: %v\n", s4)
	//当拼接操作很长时 需要换行 +号需要留在上面
	s5 := "hello" +
		"hello" +
		"hello"
	fmt.Printf("s5: %v\n", s5)
	//基本数据类型默认值

	var a int     //0
	var b float32 //0
	var c float64 //0
	var d bool    //false
	var e string  //""

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

}
