package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("The value:", m["answer"])

	m["answer"] = 48
	fmt.Println("The value:", m["answer"])

	delete(m, "answer")
	fmt.Println("The value:", m["answer"])

	value, doesExist := m["answer"]
	fmt.Println("The value:", value, "Present?", doesExist)

}
