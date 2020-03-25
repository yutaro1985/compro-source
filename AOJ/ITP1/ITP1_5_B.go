package main

import "fmt"

import "strings"

func main() {
	for {
		var H, W int
		fmt.Scan(&H, &W)
		if H == 0 {
			break
		}
		for i := 1; i <= H; i++ {
			if i == 1 || i == H {
				fmt.Println(strings.Repeat("#", W))
			} else {
				fmt.Println("#" + strings.Repeat(".", W-2) + "#")
			}
		}
		fmt.Println()
	}
}
