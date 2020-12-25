package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's saturday?")

	current := time.Now().String()

	println("current %s", current)

	today := time.Now().Weekday()

	saturday := time.Saturday

	switch saturday {
	case today + 0:
		fmt.Println("Today")

	case today + 1:
		fmt.Println("Tomorrow")

	case today + 2:
		fmt.Println("In two days")

	default:
		fmt.Println("Too far away")

	}

}
