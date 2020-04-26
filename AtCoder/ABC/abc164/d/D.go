package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	var ans int
	fmt.Scan(&S)
	for i := 4; i < len(S); i++ {
		for j := 0; i+j <= len(S); j++ {
			check, _ := strconv.Atoi(S[j : i+j])
			if check%2019 == 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
