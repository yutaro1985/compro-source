package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000009
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	n := nextInt()
	k := nextInt()
	if n <= k {
		fmt.Println(0)
		return
	}
	hn := make([]int, n)
	var totalh int
	for i := 0; i < n; i++ {
		hn[i] = nextInt()
		totalh += hn[i]
	}
	sort.Ints(hn)

	if k != 0 {
		for j := 0; j < k; j++ {
			totalh -= hn[n-j-1]
		}
	}
	fmt.Println(totalh)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
