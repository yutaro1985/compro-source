package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

// DFSの場合

func main() {
	N, M, X := nextInt(), nextInt(), nextInt()
	Cn := make([]int, N)
	Anm := make([][]int, N)
	A := make([]int, 0)
	ans := int(1e7)
	for i := range Anm {
		Anm[i] = make([]int, M)
	}
	for i := 0; i < N; i++ {
		Cn[i] = nextInt()
		for j := 0; j < M; j++ {
			Anm[i][j] = nextInt()
		}
	}
	dfs(A, Cn, Anm, N, M, X, &ans)
	if ans == int(1e7) {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func dfs(A []int, Cn []int, Anm [][]int, N, M, X int, ans *int) {
	if len(A) == N {
		check := make([]int, N)
		if !reflect.DeepEqual(A, check) {
			sum := 0
			checkUS := make([]int, M)
			for i := 0; i < N; i++ {
				if A[i] == 1 {
					sum += Cn[i]
					for j := 0; j < M; j++ {
						checkUS[j] += Anm[i][j]
					}
				}
				isOk := true
				for _, v := range checkUS {
					if v < X {
						isOk = false
					}
				}
				if isOk {
					if *ans > sum {
						*ans = sum
					}
				}
			}
		}
		return
	}

	for i := 0; i < 2; i++ {
		A = append(A, i)
		dfs(A, Cn, Anm, N, M, X, ans)
		A = A[:len(A)-1]
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
