package main

import (
	"fmt"
)

const a int = 1
const b = "bb"
const (
	c = a
	d = a + 1
	e = d + 3
	f // 和e是一样的值
	g = len("dfgfdfg")
)
const m, n = 2, "hh"

const (
	h = "A"
	i
	j = iota // 2 前面是 0和1 出现的常量次数
	k
	l = iota
)

const (
	o = iota // 重置为0
)

func main() {
	// 1 bb 1 2 5 5 7 2 hh
	fmt.Println(a, b, c, d, e, f, g, m, n)
	fmt.Println(h, i, j, k, l, o)
}
