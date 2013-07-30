// ADVANCED EXERCISE.

package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
	z := x / 2
	for i := 0; i < 10; i++ {
		z = z - ((cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2)))
	}
	return z
}

func main() {
    fmt.Println("My answer:", Cbrt(2))
    fmt.Println("2 ^ (1/3):", cmplx.Pow(2, (1 / 3.0)))
}
