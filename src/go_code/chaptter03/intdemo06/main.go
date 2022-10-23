package main

import (
	"fmt"
	"unsafe"
)

//整数的类型
//非为有符号和无符号
//int  有符号  32位系统站4字节 64位系统占8字节
//uint 无符号  32位系统站4字节 64位系统占8字节
//rune 有符号  等价int32
//byte 无符号  等价uint8

//int8  1字节    -128~127
//int16 2字节    -2^15~2^15-1
//int32 4字节    -2^33~2^31-1
//int64 8字节	 -2^63~2^63-1

// uint8  1字节    0~255
// uint16 2字节    0~2^16-1
// uint32 4字节    0~2^32-1
// uint64 8字节	  0~2^64-1
func main() {
	var a uint = 255
	fmt.Printf("a: %v\n", a)
	//整形的使用细节

	var n1 = 100000 //?n1是什么类型  默认为int
	fmt.Printf("n1: %T\n", n1)
	//如何在程序查看某个变量的占用字节大小和数据类型 (使用较多)
	var n2 int64 = 10
	fmt.Printf("n2的类型是%T所占字节: %d\n", n2, unsafe.Sizeof(n2))
	//golang程序中整型变量在使用时,遵守保小不保大的原则: 即在保证程序正确运行下 尽量使用占用空间小的数据类型
	//bit 是计算机中最小存储单位 byte:计算机中基本存储单元  1byte=8bit  1字节=8位
}
