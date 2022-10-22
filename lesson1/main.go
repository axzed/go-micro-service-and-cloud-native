package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	a, b := 1, 2
	// 调用本地函数add
	res := add(a, b)
	fmt.Println(res)
}
