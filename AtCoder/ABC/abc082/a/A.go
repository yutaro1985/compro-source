package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if (a+b)%2 != 0 {
		fmt.Println((a+b)/2 + 1)
	} else {
		fmt.Println((a + b) / 2)
	}
}
