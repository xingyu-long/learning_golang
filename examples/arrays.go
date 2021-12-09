package main

import (
	"fmt"
)

func main() {
	// var grades [4]int
	// grades[0] = 100
	grades := [3]int{90, 100, 60}
	// grades := [...]int{90, 100, 60}
	fmt.Printf("Grades: %v, %v, %v\n", grades[0], grades[1], grades[2])

	a := [...]int{1, 2, 3}
	b := a  // `b` will have a copy of `a``. New memory
	c := &a // give the address of `a` to `c`. Pointer.
	b[1] = 10
	c[1] = 5
	fmt.Println("value: ", a)
	fmt.Println("value: ", b)
	fmt.Println("value: ", c)

	d := []int{1,2,3} // slice, reference type.
	fmt.Println(d, len(d))
	fmt.Println(d, cap(d))

	e := make([]int, 3, 100) // give the capacity first.
	fmt.Println(e)
	fmt.Printf("Length: %v\n", len(e))
	fmt.Printf("Capacity: %v\n", cap(e))

	f := []int{}
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(f))
	fmt.Printf("Capacity: %v\n", cap(f))

	f = append(f, 1)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(f))
	fmt.Printf("Capacity: %v\n", cap(f))

	h := []int{1,2,3,4,5}
	fmt.Println(h)
	i := append(h[:2], h[3:]...) // it's the same as i:= append(h[:2], 3,4,5)
	fmt.Println(i)
	fmt.Println(h)
}
