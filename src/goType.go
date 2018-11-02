package main

import (
	"fmt"
	"math"
)

type (
	文本 string // 给string取一个别名 叫做 文本
)

func main() {
	var a bool
	var b string
	var c byte
	var d []byte
	var e [1]bool     // 以上都是取默认值
	var f 文本 = "我是中文" // 使用别名
	g := "哈哈"         //  使用 := 和 不规定数据类型 自动推断
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt64)
	fmt.Println(a)
	fmt.Println(b, "**")
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
