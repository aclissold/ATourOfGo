package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Sum() float64 {
    return v.X + v.Y
}

func (v *Vertex) Perimeter() float64 {
    return v.Sum() + v.Abs()
}

func main() {
    v := &Vertex{3, 4}
    fmt.Println(v.Abs())
    fmt.Println(v.Sum())
    fmt.Println(v.Perimeter())
}
