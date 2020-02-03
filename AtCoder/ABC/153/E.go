package main

import (
	"bufio"
	"os"
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

type Majic struct {
	A  int
	B  int
	ef int
}

type ByefB []Majic

func (a ByefB) Len() int           { return len(a) }
func (a ByefB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByefB) Less(i, j int) bool { return a[i].ef < a[j].ef }

func main() {
	h := nextInt()
	n := nextInt()
	col := make([][]int, n)
	for i := range col {
		a := nextInt()
		b := nextInt()
		// 消費魔力に対して最も効率のいいものを使いたい
		col[i] = []int{a, b, a / b}
	}

}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
