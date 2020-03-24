package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	// Printlnだと表記が変わって合わなくなるので明示的にフォーマットを指定する
	fmt.Printf("%.6f %.6f\n", math.Pi*r*r, 2*math.Pi*r)
}
