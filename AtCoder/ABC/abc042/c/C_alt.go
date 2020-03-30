package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, K int

	fmt.Scan(&N, &K)
	D := make([]int, K)
	for i := range D {
		fmt.Scan(&D[i])
	}
	for i := 0; i < N*10; i++ {
		ok := true
		for _, num := range strconv.Itoa(i) {
			for _, d := range D {
				if num == rune(d+'0') {
					ok = false
				}
			}
		}
		if ok && i >= N {
			fmt.Println(i)
			return
		}
	}
}
