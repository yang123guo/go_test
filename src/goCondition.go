package main

import (
	"fmt"
)

func main() {
	a := 10
	if a, b := 1, 2; a > 0 {
		// 1 2
		fmt.Println(a, b)
	}
	// 10
	fmt.Println(a)

	// if := 1; d > 0 {
	//     fmt.Println(d)
	// }
}
