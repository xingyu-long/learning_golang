package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	var a int = 32
	// *: declaring pointer
	var b *int = &a
	fmt.Println(a, *b)
	// * : dereferencing
	a = 27
	fmt.Println(a, *b)

	c := [...]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c, d, e)

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	var anotherMs *myStruct
	anotherMs = new(myStruct)
	anotherMs.foo = 33
	// it's same as 
	// (*anotherMs).foo = 33 // compiler handles it.
	fmt.Println(anotherMs.foo)


	// slice and map use `pointer`
	mapA := map[string]string {
		"foo": "bar",
		"buzz": "fizz",
	}
	mapB := mapA
	fmt.Println(mapA, mapB)
	mapB["buzz"] = "empty"
	fmt.Println(mapA, mapB)
}
