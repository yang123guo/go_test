package main

import (
	"fmt"
)

var a = ^1                  // 一元运算符
var b = 1 ^ 2               // 二元运算符
var c = 1 << 10 << 10 >> 10 // 右移 2的10次

const ( // 根据iota和右移命名进制
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
)

/*
6:      0110
11:     1011

以下四个位运算符 先转成2进制运算 最后转为10进制
6&11    0010     上下有一个为0 就为0   与操作符  最后转为10进制  -->  2
6|11    1111     上下有一个为1 就为1   或操作符                     15
6^11    1101     两位只有一个为1就为1                              13
6&^11   0100  只看下面如果为1那么改上面的取反                        4

*/

func main() {
	// -2 3
	a, b := 1, 0
	fmt.Println(a, b)
	fmt.Println(c)
	fmt.Println(6&11, 6|11, 6^11, 6&^11)
	// fmt.Println(a && b, a || b)
	fmt.Println(KB, MB, GB, TB)
	m := 1
	var p *int = &m // 有个p指向int的指针
	//0xc00006e030 取地址   1 取值
	fmt.Println(p, *p)
}
