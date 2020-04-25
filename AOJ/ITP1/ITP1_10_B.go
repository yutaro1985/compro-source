package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, C float64
	fmt.Scan(&a, &b, &C)
	rad := C * math.Pi / 180
	h := math.Sin(rad) * b
	S := a * h / 2
	L := a + b + math.Sqrt(a*a+b*b-2*a*b*math.Cos(rad))
	fmt.Printf("%f\n%f\n%f\n", S, L, h)
}
