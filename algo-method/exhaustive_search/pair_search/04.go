package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	check := make(map[string]bool)
	for i := 0; i < N; i++ {
		var S string
		fmt.Scan(&S)
		if _, e := check[S]; e {
			fmt.Println("Yes")
			return
		}
		check[S] = true
	}
	fmt.Println("No")
}
