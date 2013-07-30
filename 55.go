package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := x / 2
	aboutEqual := false
	for !aboutEqual {
		y := z - ((math.Pow(z, 2) - x) / 2 * z)
		if math.Abs(y-z) > 0.01 {
			fmt.Println(z)
			z = y
		} else {
			aboutEqual = true
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
