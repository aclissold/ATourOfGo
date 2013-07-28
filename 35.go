package main

import "code.google.com/p/go-tour/pic"

// Note: Must be copied and pasted into tour.golang.org/#35

func main() {
	pic.Show(Pic)
	pic.Show(Pic1)
	pic.Show(Pic2)
	pic.Show(Pic3)
	pic.Show(Pic4)
	pic.Show(Pic5)
}

func Pic(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)
	for y := range data {
		value := make([]uint8, dx)
		for x := range value {
			value[x] = 1 // the integer to represent
		}
		data[y] = value
	}
	return data
}

func Pic1(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)
	for y := range data {
		value := make([]uint8, dx)
		for x := range value {
			value[x] = uint8(x) ^ uint8(y) // the integer to represent
		}
		data[y] = value
	}
	return data
}

func Pic2(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)
	for y := range data {
		value := make([]uint8, dx)
		for x := range value {
			value[x] = (uint8(x) + uint8(y)/2) // the integer to represent
		}
		data[y] = value
	}
	return data
}

func Pic3(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)
	for y := range data {
		value := make([]uint8, dx)
		for x := range value {
			value[x] = uint8(x) * uint8(y) // the integer to represent
		}
		data[y] = value
	}
	return data
}

func Pic4(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)
	for y := range data {
		value := make([]uint8, dx)
		for x := range value {
			value[x] = uint8(x) + uint8(y) // the integer to represent
		}
		data[y] = value
	}
	return data
}

func Pic5(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)
	for y := range data {
		value := make([]uint8, dx)
		for x := range value {
			value[x] = uint8(x) & uint8(y) // the integer to represent
		}
		data[y] = value
	}
	return data
}
