package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/spatial"
)

// https://github.com/gonum/gonum
// https://qiita.com/po3rin/items/82c9c0195f9e3e38e2f1

func main() {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	A := mat.NewDense(3, 4, x)
	
	fmt.Println(A)
}
