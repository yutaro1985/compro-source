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

// 解説を参考
// https://youtu.be/ENSOy8u9K9I?t=1265
// 部分集合を見る場合の手法

func main() {
	N, M, X := nextInt(), nextInt(), nextInt()
	ans := 100000*12 + 5
	Cn := make([]int, N)
	Anm := make([][]int, N)
	impossible := true
	for i := range Anm {
		Anm[i] = make([]int, M)
	}
	for i := 0; i < N; i++ {
		Cn[i] = nextInt()
		for j := 0; j < M; j++ {
			Anm[i][j] = nextInt()
		}
	}
	for i := 0; i < 1<<uint(N); i++ {
		var cost int
		checkUS := make([]int, M)
		for j := 0; j < N; j++ {
			// iのj bit目が1かどうかのチェック
			// https://qiita.com/3x8tacorice/items/0b8341d7fd3ff3779111
			if uint(i)>>uint(j)&1 == 1 {
				cost += Cn[j]
				for k := 0; k < M; k++ {
					checkUS[k] += Anm[j][k]
				}
			}
		}
		ok := true
		for j := 0; j < M; j++ {
			if checkUS[j] < X {
				ok = false
			}
		}
		if ok && ans > cost {
			impossible = false
			ans = cost
		}
	}
	if impossible {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
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
