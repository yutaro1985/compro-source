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
	check := Ab.Sign() * Bb.Sign()
	if Ab.Sign() == -1 {
		Ab.Neg(Ab)
	}
	if Bb.Sign() == -1 {
		Bb.Neg(Bb)
	}
	ans := new(big.Int)
	ans = ans.Div(Ab, Bb)
	if check < 0 {
		ans.Neg(ans)
	}
	fmt.Println(ans)
}
