package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	var s string
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		num := strconv.Itoa(i)
		if i%3 == 0 {
			s += " " + num
		} else {
			for _, d := range num {
				if d == '3' {
					s += " " + num
					break
				}
			}
		}
	}
	fmt.Println(s)
}
