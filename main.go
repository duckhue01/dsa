package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	res := sum(1, 10)
	fmt.Println(res)
}
