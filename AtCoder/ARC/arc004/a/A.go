package main

import (
	"bufio"
	"fmt"
	"math"
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

// https://github.com/yutaro1985/compro-source/issues/13
func main() {
	N := nextInt()
	max := 0.0
	P := make([][]float64, N)
	for i := range P {
		P[i] = make([]float64, 2)
	}
	for i := 0; i < N; i++ {
		P[i][0], P[i][1] = float64(nextInt()), float64(nextInt())
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			l := math.Sqrt((P[j][0]-P[i][0])*(P[j][0]-P[i][0]) + (P[j][1]-P[i][1])*(P[j][1]-P[i][1]))
			if max < l {
				max = l
			}
		}
	}
	fmt.Println(max)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
