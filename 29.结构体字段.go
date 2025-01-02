package main

import "fmt"

type Vertex1 struct {
	X int
	Y int
}

func main() {
	v := Vertex1{11, 2}
	v.X = 4
	fmt.Println(v.X)
}
