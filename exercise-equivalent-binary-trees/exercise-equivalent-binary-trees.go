package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walk(t, ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }