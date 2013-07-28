package main

import "fmt"

func main() {
	var z []int
	fmt.Printf("%T\n", z)
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
