package main

import (
	"bufio"
	"fmt"
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
	E := make([]int, 6)
	L := make(map[int]int)
	for i := range E {
		E[i] = nextInt()
	}
	B := nextInt()
	for i := 0; i <= 5; i++ {
		L[nextInt()] = 0
	}
	var hit int
	var bonus bool
	if _, ok := L[B]; ok {
		bonus = true
	} else {
		bonus = false
	}
	for _, Ei := range E {
		if _, ok := L[Ei]; ok {
			hit++
		}
	}
	ans := 0
	switch hit {
	case 6:
		ans = 1
	case 5:
		if bonus {
			ans = 2
		} else {
			ans = 3
		}
	case 4:
		ans = 4
	case 3:
		ans = 5
	default:
		ans = 0
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
