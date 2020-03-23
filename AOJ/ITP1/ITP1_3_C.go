package main

import "fmt"

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)
		if x == 0 && y == 0 {
			break
		} else {
			if x < y {
				fmt.Println(x, y)
			} else {
				fmt.Println(y, x)
			}
		}
	}
}
