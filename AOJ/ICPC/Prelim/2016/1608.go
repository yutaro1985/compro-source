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

func main() {
	for {
		n := nextInt()
		if n == 0 {
			break
		}
		an := make([]int, n)
		min := 1000000
		for i := 0; i < n; i++ {
			an[i] = nextInt()
		}
		sort.Ints(an)
		for i := 1; i < n; i++ {
			sub := an[i] - an[i-1]
			if min > sub {
				min = sub
			}
		}
		fmt.Println(min)
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
