package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(min(min(abs(a, b), abs(a+10, b)), abs(a, b+10)))
}

func abs(a int, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
