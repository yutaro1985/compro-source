package main

import "fmt"

import "strconv"

func main() {
	for i := 1; ; i++ {
		var x int
		fmt.Scan(&x)
		if x != 0 {
			fmt.Println("Case " + strconv.Itoa(i) + ": " + strconv.Itoa(x))
		} else {
			break
		}
	}
}
