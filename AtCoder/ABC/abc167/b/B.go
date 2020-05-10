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
	A, B, C, K := nextInt(), nextInt(), nextInt(), nextInt()
	var ans int
	if A > K {
		fmt.Println(K)
		return
	}
	ans += A
	K -= A
	K -= B
	if K <= 0 {
		fmt.Println(ans)
		return
	}
	if K > 0 {
		if K > C {
			ans -= C
		} else {
			ans -= K
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
