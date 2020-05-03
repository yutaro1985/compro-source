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
	N, M := nextInt(), nextInt()
	H := make([]int, N)
	roads := make([][]int, N)
	var ans int
	for i := range H {
		H[i] = nextInt()
	}
	for i := 0; i < M; i++ {
		A, B := nextInt()-1, nextInt()-1
		roads[A] = append(roads[A], B)
		roads[B] = append(roads[B], A)
	}
	for i := 0; i < N; i++ {
		if len(roads[i]) != 0 {
			good := true
			for _, dist := range roads[i] {
				if H[dist] >= H[i] {
					good = false
				}
			}
			if good {
				ans++
			}
		} else {
			ans++
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
