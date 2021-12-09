package main

import (
	"fmt"
)

const a int16 = 27

const (
	i1 = iota // 0
	i2 = iota // 1
	i3 = iota // 2
)


const (
	_ = iota
	KB = 1 << (10 * iota)
	MB 
	GB
	TB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func main() {
	const myConst int = 5
	fmt.Printf("%v %T\n", myConst, myConst)
	const a int16 = 123
	fmt.Printf("%v %T\n", a, a)

	const c = 42
	var d int16 = 27
	fmt.Printf("%v %T\n", c+d, c+d)

	fmt.Printf("%v %T\n", i1, i1)
	fmt.Printf("%v %T\n", i2, i2)
	fmt.Printf("%v %T\n", i3, i3)

	fmt.Printf("%v %T\n", KB, KB)
	fmt.Printf("%v %T\n", MB, MB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Can see Asia? %v\n", canSeeAsia & roles == canSeeAsia)
}
