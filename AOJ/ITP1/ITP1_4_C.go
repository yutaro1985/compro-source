package main

import "fmt"

func main() {
	for {
		var a, b int
		var op string
		fmt.Scan(&a, &op, &b)
		switch op {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		default:
			return
		}
	}
}
