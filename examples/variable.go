package main

import (
	"fmt"
	"strconv"
)

var (
	actorName   string = "A"
	actorGender string = "B"
)

// variable scope:
// 	package: lowercase, var i int = 42, available to the package
// 	global: uppercase, var I int = 42, globally visible.
//  block: inside of function.
func main() {
	// variable declaration
	var i int
	i = 21
	var j int = 28
	k := 42
	fmt.Println(i, j, k)

	// explicit conversion
	var w float32
	w = float32(i)
	fmt.Printf("%v %T\n", w, w)

	var s string
	s = strconv.Itoa(i) // use strconv package for strings.
	fmt.Printf("%v %T\n", s, s)
}
