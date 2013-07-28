package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Printf("%T %v\n", v, v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}
