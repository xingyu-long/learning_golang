package main

import (
	"fmt"
)

func main()  {
	// defer, panic, recover

	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")


	a := "start"
	defer fmt.Println(a) // use `a` when it was called.
	a = "end"

	c, d := 1, 0
	res := c / d
	fmt.Println(res)
	// panic("msg")
	// we can use panic for some errors.
}