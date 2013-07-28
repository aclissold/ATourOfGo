package main

import "fmt"

type vertex struct {
	X, Y int
}

var (
	p = vertex{1, 2}  // has type Vertext
	q = &vertex{1, 2} // has type *Vertext
	r = vertex{X: 1}  // Y:0 is implicit
	s = vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
