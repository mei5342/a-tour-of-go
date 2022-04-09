package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printlSlice(s)

	s = s[:4]
	printSlice(s)
}