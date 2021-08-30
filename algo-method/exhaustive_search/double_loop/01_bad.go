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

var buf []byte = make([]byte, initialBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	N := nextInt()
	isPrime := map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true, 23: true, 29: true, 31: true, 37: true, 41: true, 43: true, 47: true, 53: true, 59: true, 61: true, 67: true, 71: true, 73: true, 79: true, 83: true, 89: true, 97: true,
	}
	var ans int
	for i := 0; i < N; i++ {
		if isPrime[nextInt()] {
			ans++
		}
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
