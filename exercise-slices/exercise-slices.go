package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pow := make([][]uint8, dy)
	for y := range pow {
		pow[y] = make([]uint8, dx)
		for x := range pow[y] {
			pow[y][x] = uint8((x + y) / 2)
		}
	}
	return pow
}

func main () {
	pic.Show(Pic)
}