package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func main() {

	var (
		a int
		b int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println((a + b) * (a + b))
}
