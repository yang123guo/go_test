package main

import (
	"fmt"
	"math"
	"strconv"
)

type (
	文本 string // 给string取一个别名 叫做 文本
)

func main() {
	var a bool
	var b string
	var c byte
	var d []byte
	var e [1]bool            // 以上都是取默认值
	var f 文本 = "我是中文"        // 使用别名
	g := "哈哈"                //  使用 := 和 不规定数据类型 自动推断
	h, _, i, j := 1, 2, 3, 4 // _代表的空位符号
	var m float32 = 100.1
	n := int(m)
	var r bool = true
	var t int
	var s int = 65
	w := string(s)
	v := strconv.Itoa(s) // 转化为字符串 65
	// t := int(r)  // bool不能转为int或者float
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt64)
	fmt.Println(a)
	fmt.Println(b, "**")
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h, i, j)
	fmt.Println(m, n, r, t)
	fmt.Println(s, w, v)
}
