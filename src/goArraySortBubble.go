package main

import "fmt"

func main() {
	a := [...]int{5, 3, 5, 7, 2, 8, 11}
	fmt.Println(a) // [5 3 5 7 2 8 11]

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ { // i和j(i+1)比较
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a) // [11 8 7 5 5 3 2]
}
