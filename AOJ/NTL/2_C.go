package main

import (
	"fmt"
	"math/big"
)

func main() {
	var A, B string
	Ab, Bb := new(big.Int), new(big.Int)
	fmt.Scan(&A, &B)
	_, _ = fmt.Sscan(A, Ab)
	_, _ = fmt.Sscan(B, Bb)
	ans := Ab.Mul(Ab, Bb)
	fmt.Println(ans)
}
