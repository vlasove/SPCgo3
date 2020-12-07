//main.go
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
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	result := add(a, b)*sub(b, a) - mult(a, a)
	fmt.Println("Result:", result)
}
