package main

import (
	"fmt"
)

func main() {
	var S, T, U string
	var A, B int
	fmt.Scan(&S, &T)
	fmt.Scan(&A, &B)
	fmt.Scan(&U)
	if S == U {
		fmt.Println(A-1, B)
	} else {
		fmt.Println(A, B-1)
	}
}
