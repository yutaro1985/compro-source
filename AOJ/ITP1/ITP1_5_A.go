package main

import "fmt"

func main() {
	for {
		var H, W int
		var col string
		fmt.Scan(&H, &W)
		if H == 0 {
			break
		}
		for i := 0; i < W; i++ {
			col += "#"
		}
		for i := 0; i < H; i++ {
			fmt.Println(col)
		}
		fmt.Println()
	}
}
