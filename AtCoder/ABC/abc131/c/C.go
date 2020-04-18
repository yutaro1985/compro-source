package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	A, B, C, D := nextInt(), nextInt(), nextInt(), nextInt()
	var divAC, divAD, divBC, divBD, divACD, divBCD int
	if A%C == 0 {
		// A未満にしたい
		divAC = A/C - 1
	} else {
		divAC = A / C
	}
	if A%D == 0 {
		divAD = A/D - 1
	} else {
		divAD = A / D
	}
	lcm := lcmof2numbers(C, D)
	if A%lcm == 0 {
		divACD = A/lcm - 1
	} else {
		divACD = A / lcm
	}
	divBCD = B / lcm
	// こちらはB以下でいい
	divBC = B / C
	divBD = B / D
	Cdivs := divBC - divAC
	Ddivs := divBD - divAD
	CDdivs := divBCD - divACD
	fmt.Println(B - A + 1 - (Cdivs + Ddivs - CDdivs))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}

func lcmof2numbers(a int, b int) int {
	return a / gcdof2numbers(a, b) * b
}
