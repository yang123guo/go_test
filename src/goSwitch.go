package main

import (
	"fmt"
)

func main() {
	a := 1
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}

	b := 2
	switch {
	case b > 0:
		fmt.Println("0000")
		fallthrough
	case b > 1:
		fmt.Println("11111")
	}

	switch c := 2; {
	case c > 0:
		fmt.Println("ccc")
		fallthrough
	case c > 1:
		fmt.Println("111ccc11")
	}
}
