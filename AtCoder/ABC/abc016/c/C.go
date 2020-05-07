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
	friends := make([]map[int]bool, N)
	fcount := make([]map[int]bool, N)
	for i := 0; i < N; i++ {
		friends[i] = make(map[int]bool)
		fcount[i] = make(map[int]bool)
	}
	for i := 0; i < M; i++ {
		A, B := nextInt()-1, nextInt()-1
		friends[A][B] = true
		friends[B][A] = true
	}
	for i := 0; i < N; i++ {
		for k := range friends[i] {
			for k2 := range friends[k] {
				if _, exist := friends[i][k2]; !exist && k2 != i {
					fcount[i][k2] = true
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		fmt.Println(len(fcount[i]))
	}
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
