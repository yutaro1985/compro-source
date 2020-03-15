package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Go 1.6ではbig.SqrtがまだないのでCEになる
	// https://github.com/golang/go/commits/master/src/math/big/sqrt.go
	var a, b, c big.Float
	var as, bs, cs, sum big.Float
	fmt.Scan(&a, &b, &c)
	sum.SetPrec(1024)
	as.Sqrt(&a)
	bs.Sqrt(&b)
	cs.Sqrt(&c)
	sum.Add(&as, &bs)
	if cs.Cmp(&sum) == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
