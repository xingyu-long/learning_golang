package main

import (
	"fmt"
)

func main() {
	var n bool = false
	fmt.Printf("%v %T\n", n, n)
	a := 10 // 1010
	b := 3  // 0011
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)

	var complex complex64 = 1 + 2i
	fmt.Printf("%v %T\n", real(complex), real(complex))
	fmt.Printf("%v %T\n", imag(complex), imag(complex))

	s := "this is a string"
	byteOfs := []byte(s)
	fmt.Printf("%v %T\n", byteOfs, byteOfs)

	// r := 'a'
	// rune: Alias int32
	var r rune = 'a'
	fmt.Printf("%v %T\n", r, r)
}
