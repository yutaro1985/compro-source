package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	N := nextInt()
	An := make([]int, N)
	Bn := make([]int, N)
	var sumA, MaxB int
	ans := "Yes"
	for i := 0; i < N; i++ {
		An[i], Bn[i] = nextInt(), nextInt()
		sumA += An[i]
		if MaxB < Bn[i] {
			MaxB = Bn[i]
		}
	}
	sort.Ints(An)
	sort.Ints(Bn)
	for i := N - 1; i >= 0; i-- {
		// fmt.Println(sumA, An[i], Bn[i])
		if sumA > Bn[i] {
			ans = "No"
			break
		}
		sumA -= An[i]
	}

	fmt.Println(ans)
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
