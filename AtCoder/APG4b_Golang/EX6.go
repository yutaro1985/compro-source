package main

import "fmt"

func main() {
	var A, B int
	var S string
	fmt.Scan(&A, &S, &B)
	switch S {
	case "+":
		fmt.Println(A + B)
	case "-":
		fmt.Println(A - B)
	case "*":
		fmt.Println(A * B)
	case "/":
		if B != 0 {
			fmt.Println(A / B)
		} else {
			fmt.Println("error")
		}
	default:
		fmt.Println("error")
	}
}
