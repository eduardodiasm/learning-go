package main

import "fmt"

type Vertex struct {
	X int64
	Y int64
}

func main() {
	v := Vertex{1, 2}
	v.X = 4

	fmt.Println(v.X)

}
