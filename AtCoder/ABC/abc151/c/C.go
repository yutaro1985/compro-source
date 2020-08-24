package main

import (
	"bufio"
	"fmt"
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

func main() {
	n := nextInt()
	m := nextInt()
	clearedlist := make([]bool, n)
	penaltycount := make([]int, n)
	clearcount := 0
	penalty := 0

	for i := 0; i < m; i++ {
		p := nextInt()
		s := nextLine()
		if clearedlist[p-1] == true {
			continue
		}
		if s == "AC" {
			clearedlist[p-1] = true
			clearcount++
		} else if s == "WA" {
			penaltycount[p-1]++
		}
	}
	for i, p := range penaltycount {
		if clearedlist[i] == true {
			penalty += p
		}
	}
	fmt.Println(clearcount, penalty)
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
