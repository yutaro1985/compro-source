package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	condition := (c-a-b)*(c-a-b) - 4*a*b
	if condition > 0 && c > a+b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
