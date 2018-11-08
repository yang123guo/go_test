package main

import "fmt"

func main() {
	var slice1 []int  // 空slice
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 长度为10的数组
	a9 = a[9]
	slice2 := a[5:10] // 相当于a[56789] 左闭右开
	slice3 := a[5:]  // 从索引5到尾部
	slice4 := a[:5] // 去首部到前五个 0-4

	// 另外一种申明方法
	// 如果不设置容量 那么容量就是3
	slice5 := make([]int, 3, 10) // int 类型 3元素个数len  容量(最大10，可以小于等于) 超出以后会扩容 系统重新分配，但是效率低
	fmt.Println(slice1)
	fmt.Println(len(slice5), cap(slice5))
}
