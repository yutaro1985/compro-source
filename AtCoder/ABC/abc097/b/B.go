package main

import (
	"fmt"
	"math"
)

func main() {
	var X int
	fmt.Scan(&X)
	ans := 1
	for i := 2; i <= X; i++ {
		for j := 2; j <= X; j++ {
			num := int(math.Pow(float64(i), float64(j)))
			if num > X {
				break
			} else if num > ans {
				ans = num
			}
		}
	}
	fmt.Println(ans)
}
