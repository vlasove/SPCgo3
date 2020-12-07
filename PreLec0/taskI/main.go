//TPP:I solution
package main

import "fmt"

func main() {

	var A, B, C int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	if A == 0 {
		//Bx + C  = 0
		//x = -C/B
		if B == 0 {
			fmt.Println("НЕТ КОРНЕЙ")
		} else {
			fmt.Println("ОДИН КОРЕНЬ")
		}
	} else {
		//Ax^2 + Bx + C = 0
		D := B*B - 4*A*C
		if D > 0 {
			fmt.Println("ДВА КОРНЯ")
		} else if D == 0 {
			fmt.Println("ОДИН КОРЕНЬ")
		} else {
			fmt.Println("НЕТ КОРНЕЙ")
		}
	}
}
