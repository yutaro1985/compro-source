package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		fmt.Printf("%d", a)
	} else if a == b {
		fmt.Printf("eq")
	} else {
		fmt.Printf("%d", b)
	}
}
