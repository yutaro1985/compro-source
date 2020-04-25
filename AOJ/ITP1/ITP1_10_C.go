package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var n, a, sum float64
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		s := make([]float64, int(n))
		for i := range s {
			fmt.Scan(&s[i])
			sum += s[i]
		}
		m := sum / n
		for _, si := range s {
			a += (si - m) * (si - m)
		}
		fmt.Println(math.Sqrt(a / n))
	}
}
