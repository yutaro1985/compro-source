package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	if B > A*C {
		fmt.Println(C)
	} else {
		fmt.Println(B / A)
	}
}
