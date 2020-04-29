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
	ans := "YES"
	list := map[int]bool{}
	for i := 0; i < 10; i++ {
		list[i] = false
	}
	check := []int{1, 9, 7, 4}
	for i := 0; i < 4; i++ {
		list[nextInt()] = true
	}
	for _, c := range check {
		if !list[c] {
			ans = "NO"
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
