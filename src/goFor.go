package main

import (
	"fmt"
)

func main() {
	a := 1

	for {
		a++
		if a > 4 {
			break
		}
		fmt.Println(a, "a")
	}

	b := 1
	for b < 3 {
		b++
		fmt.Println(b, "b")
	}

	for c := 0; c < 4; c++ {
		fmt.Println(c, "c")
	}

}
