package main

import "fmt"

func main() {
	// 呼び出した関数を関数の終わりまで遅延させる
	defer fmt.Println("world")

	fmt.Println("hello")
}