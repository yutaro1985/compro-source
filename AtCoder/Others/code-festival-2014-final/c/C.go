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

func main() {
	A := nextInt()
	for i := 10; i <= 10000; i++ {
		ans, cnt := 0, 0
		for j := i; j != 0; j /= 10 {
			ans += (j % 10) * int(math.Pow(float64(i), float64(cnt)))
			cnt++
		}
		if ans == A {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
