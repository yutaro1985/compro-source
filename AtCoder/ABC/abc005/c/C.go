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
	T, N := nextInt(), nextInt()
	An := make([]int, N)
	for i := range An {
		An[i] = nextInt()
	}
	M := nextInt()
	Bn := make([]int, M)
	for i := range Bn {
		Bn[i] = nextInt()
	}
	ans := "no"
	j := 0
	for i := 0; i < N; i++ {
		if An[i]+T >= Bn[j] && An[i] <= Bn[j] {
			j++
		}
		if j == len(Bn) {
			ans = "yes"
			break
		}
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
