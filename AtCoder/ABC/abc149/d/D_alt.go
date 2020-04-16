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
	ans := 0
	v := map[byte]int{
		'r': P,
		'p': S,
		's': R,
	}
	pat := make([][]byte, K)
	for i := 0; i < N; i++ {
		pat[i%K] = append(pat[i%K], T[i])
	}
	for i := 0; i < K; i++ {
		ans += v[T[i]]
		w := true
		for j := 1; j < len(pat[i]); j++ {
			if !w {
				ans += v[pat[i][j]]
				w = true
			} else {
				if pat[i][j] != pat[i][j-1] {
					ans += v[pat[i][j]]
					w = true
				} else {
					w = false
				}
			}
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
