package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		var s string
		check := make([]string, 5)
		if i%2 == 0 {
			check[0] = "a"
		}
		if i%3 == 0 {
			check[1] = "b"
		}
		if i%4 == 0 {
			check[2] = "c"
		}
		if i%5 == 0 {
			check[3] = "d"
		}
		if i%6 == 0 {
			check[4] = "e"
		}
		for _, letter := range check {
			s += letter
		}
		if s != "" {
			fmt.Println(s)
		} else {
			fmt.Println(i)
		}

	}
}
