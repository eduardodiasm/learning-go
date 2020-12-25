package main

import "fmt"

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

// Multiple returns
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

}
