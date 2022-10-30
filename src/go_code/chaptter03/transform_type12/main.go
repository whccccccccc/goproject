package main

import "fmt"

//演示基本数据类转换
func main() {
	var i int = 100
	//int 转换float
	var f1 float32 = float32(i)
	fmt.Printf("f1: %v\n", f1)
	fmt.Printf("f1: %T\n", f1)
	//转回int
	var i1 = int32(f1)
	fmt.Printf("i1: %T\n", i1)
	//转换int64
	var i2 = int64(i1)
	fmt.Printf("i2: %T\n", i2)
	//细节说明
	//1.数据类的转换可以从范围小->范围大 也可以范围大->范围小
	//2.被转换的时变量存储的数据(即值) 变量本身的数据类型并没有变化  创建了一份新的变量出来
	//3.在转换中 大转小  可能出现溢出问题
	var i3 = int32(i1)
	fmt.Printf("i3: %T\n", i3)
	var i4 int32 = 255555555
	var i5 = int8(i4)
	fmt.Printf("i4: %v\n", i4)
	fmt.Printf("i5: %v\n", i5)

}
