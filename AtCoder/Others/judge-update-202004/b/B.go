package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// go 1.14.1
const (
	initialBufSize = 1e4
	maxBufSize     = 1e8 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

type bag struct {
	number int
	color  string
}

type ByColor []bag

func (a ByColor) Len() int           { return len(a) }
func (a ByColor) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByColor) Less(i, j int) bool { return a[i].color < a[j].color }

func main() {
	b := []bag{}
	N := nextInt()
	for i := 0; i < N; i++ {
		b = append(b, bag{nextInt(), nextLine()})
	}
	// go 1.6だとできないソート
	// https://c2masamichi.hatenablog.com/entry/2019/12/16/225519
	// 複数キーのソート
	// https://or3.hatenablog.com/entry/2018/01/15/212029
	sort.Slice(b, func(i, j int) bool {
		if b[i].color > b[j].color {
			return true
		} else if b[i].color < b[j].color {
			return false
		} else {
			return b[i].number < b[j].number
		}
	})
	for _, bag := range b {
		fmt.Println(bag.number)
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
