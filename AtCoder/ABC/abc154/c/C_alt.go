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
	// 解説を参考にした別解
	// > そのため、一度与えられた数列をソートすると、全ての隣接する 2 項を比較し、
	// > 等しい値がなければ YES で、等しい値があれば NO です。
	N := nextInt()
	list := make([]int, N)
	for i := 0; i < N; i++ {
		list[i] = nextInt()
	}
	sort.Ints(list)
	for i := 1; i < N; i++ {
		if list[i-1] == list[i] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
