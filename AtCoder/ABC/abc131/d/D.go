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

type limit struct {
	A int
	B int
}

type sortB []limit

func (a sortB) Len() int           { return len(a) }
func (a sortB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortB) Less(i, j int) bool { return a[i].B < a[j].B }

func main() {
	N := nextInt()
	ABn := []limit{}
	for i := 0; i < N; i++ {
		ABn = append(ABn, limit{nextInt(), nextInt()})
	}
	sort.Sort(sortB(ABn))
	ans := "Yes"
	sumA := 0
	for i := 0; i < N; i++ {
		sumA += ABn[i].A
		if ABn[i].B < sumA {
			ans = "No"
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
