package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

// method for `greeter` struct.
// `g` is a copy of greeter struct.
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// pointer receiver
// access same with address.
func (g *greeter) greetWithChange() {
	// implicit dereferencing.
	fmt.Println(g.greeting, g.name)
	g.name = "Changed"
}

func main() {
	sayMessage("Hello Go!")
	sayGreeting("Hello", "World")
	sum([]int{1, 2, 3, 4, 5, 6}...)
	sum(1, 2, 3, 4, 5, 6)

	// Anonymous function
	func() {
		fmt.Println("Here is an Anonymous function")
	}()

	g := greeter{
		greeting: "Hello",
		name:     "Go from greeter",
	}
	g.greet()
	fmt.Printf("name = %v\n", g.name)
	g.greetWithChange()
	fmt.Printf("name = %v\n", g.name)

	// return with err.
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}

func sayMessage(msg string) {
	fmt.Println(msg)
}

// it's same as (greeting string, name string)
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

// ...int should be placed at end.
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Printf("Sum = %v\n", result)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
