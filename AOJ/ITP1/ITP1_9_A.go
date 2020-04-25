package main

import (
	"fmt"
	"strings"
)

func main() {
	var W string
	var ans int
	fmt.Scan(&W)
	for {
		var T string
		fmt.Scan(&T)
		if strings.ToLower(T) == W {
			ans++
		}
		if T == "END_OF_TEXT" {
			break
		}
	}
	fmt.Println(ans)
}
