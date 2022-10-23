package main

import "fmt"

//小数类型的使用
//小数类型分类
//单精度float32  4字节
//双精度float64  8字节
func main() {
	var price float32 = 89.12
	fmt.Printf("price: %v\n", price)
	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Printf("num1: %v\n", num1)
	fmt.Printf("num2: %v\n", num2)
	//尾数部分可能丢失 造成精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901 //高精度用这个
	fmt.Printf("num3: %v\n", num3)
	fmt.Printf("num4: %v\n", num4)
	//1.Golang的浮点类型有固定的范围和字段长度不受具体操作系统的影响
	//2.Golang的浮点型默认声明为 float64类型
	//3.表现形式有两种
	num6 := 5.12
	num7 := .123 //不建议这样写 带上0比较好
	fmt.Printf("num6: %v\n", num6)
	fmt.Printf("num7: %v\n", num7)
	//科学计数法形式
	num8 := 5.1234e2   //  =5.1234*10的2次方
	num9 := 5.1234e2   //  =5.1234*10的2次方
	num10 := 5.1234e-2 //  =5.1234 / 10的2次方
	fmt.Printf("num8: %v\n", num8)
	fmt.Printf("num9: %v\n", num9)
	fmt.Printf("num10: %v\n", num10)
	//4.通常情况下应该使用float64 因为更精确
}
