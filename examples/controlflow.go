package main

import (
	"fmt"
)

func main() {
	// ||, &&, !
	if true {
		fmt.Println("The test is True")
	}

	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else {
		fmt.Println("Just right here!")
	}

	switch i := 2 + 3; i {
	case 1, 3, 5, 7:
		fmt.Println("odd")
	case 2, 4, 6, 8:
		fmt.Println("even")
	default:
		fmt.Println("not one or two")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // make next one working.
	default:
		fmt.Println("greater than twenty")
	}

	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("i is am int.")
		// break // you can use break to terminate earlier.
	case float64:
		fmt.Println("i is a float64.")
	default:
		fmt.Println("i is another type")
	}
}
