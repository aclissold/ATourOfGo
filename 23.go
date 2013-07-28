// TODO: print out how many iterations to == math.Sqrt
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
    z := x / 2
    aboutEqual := false
    for !aboutEqual {
        y := z - ((math.Pow(z, 2) - x) / 2 * z)
        if math.Abs(y - z) > 0.001 {
            fmt.Println(z)
            z = y
        } else {
            aboutEqual = true
        }
    }
    return z
}

func main() {
    const check = 2
    fmt.Println("My approximation:")
	fmt.Println(Sqrt(check))
    fmt.Println("math.Sqrt's:")
    fmt.Println(math.Sqrt(check))
}
