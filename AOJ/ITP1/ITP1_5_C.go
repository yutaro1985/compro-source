package main

import "fmt"

import "strings"

func main() {
	for {
		var H, W int
		var line string
		fmt.Scan(&H, &W)
		if H == 0 {
			break
		}
		for i := 0; i < H; i++ {
			// こちらのほうがスマートな書き方
			// for j := 0; j < W; j++ {
			// 	if (i+j)%2 != 0 {
			// 		fmt.Print("#")
			// 	} else {
			// 		fmt.Print(".")
			// 	}
			// }
			// fmt.Println()
			if i%2 != 0 {
				if W%2 != 0 {
					line = strings.Repeat(".#", W/2) + "."
				} else {
					line = strings.Repeat(".#", W/2)
				}
			} else {
				if W%2 != 0 {
					line = strings.Repeat("#.", W/2) + "#"
				} else {
					line = strings.Repeat("#.", W/2)
				}
			}
			fmt.Println(line)
		}
		fmt.Println()
	}
}
