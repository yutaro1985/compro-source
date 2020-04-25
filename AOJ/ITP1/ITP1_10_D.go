package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	xn := make([]float64, n)
	yn := make([]float64, n)
	for i := range xn {
		fmt.Scan(&xn[i])
	}
	for i := range yn {
		fmt.Scan(&yn[i])
	}
	var D1, D2, D3, Di float64
	for i := 0; i < n; i++ {
		d := xn[i] - yn[i]
		D1 += math.Abs(d)
		D2 += d * d
		D3 += math.Abs(d * d * d)
		if Di < math.Abs(d) {
			Di = math.Abs(d)
		}
	}
	fmt.Printf("%f\n%f\n%f\n%f\n", D1, math.Sqrt(D2), math.Cbrt(D3), Di)
}
