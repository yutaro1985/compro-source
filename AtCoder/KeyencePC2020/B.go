package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
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
	N := nextInt()
	m := make(map[int]int)
	keys := make([]int, 0, N)
	max := -1000000009
	for i := 0; i < N; i++ {
		x := nextInt()
		m[x] = nextInt()
	}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i := 1; i < N; i++ {
		if keys[i-1] + m[keys[i-1]]  >= max {
			max = keys[i-1] + m[keys[i-1]]
		}
		if max > keys[i] - m[keys[i]] {
			delete(m,keys[i])
		}
	}
	fmt.Println(len(m))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
