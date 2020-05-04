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
	N := nextInt()
	var sum int
	city := map[string]int{}
	for i := 0; i < N; i++ {
		name := nextLine()
		population := nextInt()
		city[name] = population
		sum += population
	}
	for name, population := range city {
		if population > sum/2 {
			fmt.Println(name)
			return
		}
	}
	fmt.Println("atcoder")
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
