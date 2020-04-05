package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if abs(a, c) <= d || (abs(a, b) <= d && abs(b, c) <= d) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
