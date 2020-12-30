package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var myMap map[string]Vertex

func main() {

	myMap = make(map[string]Vertex)

	myMap["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(myMap["Bell Labs"])

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)

	var mm = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(mm)

}
