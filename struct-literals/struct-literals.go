package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &Vertex{1, 2} // & を頭に付けると、新しく割り当てられたstructへのポインタを戻す。
)

func main() {
	fmt.Println(v1, p, v2, v3)
}