package main

import "fmt"

func fibonacci() func() int {
	f0, f1 := 0, 1
	return func() int {
		v := f0
		f0, f1 = f1, f0 + f1
		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}