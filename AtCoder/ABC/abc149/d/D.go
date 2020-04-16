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
	N, K, R, S, P := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	T := nextLine()
	W := make([]bool, N)
	ans := 0
	v := map[byte]int{
		'r': P,
		'p': S,
		's': R,
	}

	for i := 0; i < N; i++ {
		if i < K {
			ans += v[T[i]]
		} else if T[i] == T[i-K] {
			if W[i-K] {
				ans += v[T[i]]
			} else {
				W[i] = true
			}
		} else {
			ans += v[T[i]]
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
