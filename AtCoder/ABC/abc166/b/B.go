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
	N, K := nextInt(), nextInt()
	snuke := make([]int, N)
	for i := 0; i < K; i++ {
		d := nextInt()
		for i := 0; i < d; i++ {
			snuke[nextInt()-1]++
		}
	}
	var ans int
	for _, a := range snuke {
		if a == 0 {
			ans++
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
