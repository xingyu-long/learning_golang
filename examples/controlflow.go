package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The test is True")
	}

	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	} 
	if guess > number {
		fmt.Println("Too high")
	}
}
