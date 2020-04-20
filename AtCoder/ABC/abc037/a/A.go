package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	if A > B {
		fmt.Println(C / B)
	} else {
		fmt.Println(C / A)
	}
}
