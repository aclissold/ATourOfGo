package main

import (
	"fmt"
	"math"
)

func main() {
    x, err := fmt.Println("Happy", math.Pi, "day!")
    if err != nil {
        fmt.Println(x, err)
    }
}
