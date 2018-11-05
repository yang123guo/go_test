package main

import "fmt"

func main() {
	var (
		name   = [5]int{1, 4: 223}
		name1  = [...]float32{10: 10.1}
		name2  = [...]float32{10: 10.1}
		value1 = [10]int{}
	)
	x, y := 1, 2
	p1 := [...]*int{&x, &y}
	name2[1] = 58.0
	value2 := new([10]int)
	value1[1] = 11
	value2[1] = 11
	b := [2][3]int{{1, 2}, {2: 2}}
	fmt.Println(name)           // [1 0 58 0 223]
	fmt.Println(value1, value2) // [0 11 0 0 0 0 0 0 0 0]   &[0 11 0 0 0 0 0 0 0 0]
	fmt.Println(name1)          // [0 0 0 0 0 0 0 0 0 0 10.1]
	fmt.Println(p1)             // [0xc00008c000 0xc00008c008]
	fmt.Println(name2 == name1) // true
	fmt.Println(b)              // [[1 2 0] [0 0 2]]
}
