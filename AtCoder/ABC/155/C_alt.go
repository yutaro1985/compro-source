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
	N := nextInt()
	s := make([]string, N)
	m := make(map[string]int)
	max := 0
	for i := 0; i < N; i++ {
		s[i] = nextLine()
	}
	sort.Strings(s)
	for i := 0; i < N; i++ {
		m[s[i]]++
		if max < m[s[i]] {
			max = m[s[i]]
		}
	}
	// GolangのMapはfor rangeでの順序が保証できないので、keyを格納してsortしたスライスを用意してそれを使用する
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if m[k] == max {
			fmt.Println(k)
		}
	}
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
