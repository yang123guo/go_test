package main

import "fmt"

// 可见 *p  去的是变量的值    &a 取得指针的地址

func main() {
	var a int = 20                   /* 声明实际变量 */
	var ip *int                      /* 声明指针变量 */
	ip = &a                          /* 指针变量的存储地址 */
	fmt.Printf("a 变量的地址是: %x\n", &a) // a 变量的地址是: 20818a220
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip) // ip 变量储存的指针地址: 20818a220
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip) // *ip 变量的值: 20
}
