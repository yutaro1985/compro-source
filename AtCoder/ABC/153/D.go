package main

import (
	"fmt"
	"math"
)

func main() {
	var h int
	fmt.Scan(&h)
	total := 0
	count := 0
	if h == 1 {
		fmt.Println(1)
		return
	}
	for h >= 1 {
		total += int(math.Pow(2, float64(count)))
		if h%2 != 0 {
			h--
		}
		h /= 2
		count++
	}
	fmt.Println(total)
}
