package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a == b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
