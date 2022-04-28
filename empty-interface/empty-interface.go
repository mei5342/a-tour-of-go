package main

import "fmt"

func main() {
	var i interface{} // 空のインターフェスは任意の型の値を保持できる
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}