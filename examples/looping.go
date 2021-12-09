package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j :=0, 0; i < 5; i, j = i + 1, j + 2 {
		fmt.Println(i, j)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	// infinite loop
	for true {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	s := []int{1,2,3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
