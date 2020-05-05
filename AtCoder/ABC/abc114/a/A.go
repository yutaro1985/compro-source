package main

import "fmt"

func main() {
	var X int
	fmt.Scan(&X)
	switch X {
	case 7, 5, 3:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
