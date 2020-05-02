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

// 深さ優先探索の問題
// TODO 深さ優先探索の練習をした後に実装する

type nums struct {
	A int
	B int
	C int
	D int
}

func main() {
	N, M, Q := nextInt(), nextInt(), nextInt()
	nums := make([]nums, Q)
	for i := range nums {
		nums[i].A = nextInt() - 1
		nums[i].B = nextInt() - 1
		nums[i].C = nextInt()
		nums[i].D = nextInt()
	}
	A := make([]int, N)
	var dfs func(index int, last int) int
	dfs = func(index int, last int) int {
		if index > N {
			sum := 0
			for i := 0; i < Q; i++ {
				if A[nums[i].B]-A[nums[i].A] == nums[i].C {
					sum += nums[i].D
				}
			}
			return sum
		}
		max := 0
		for i := 0; i+last < M; i++ {
			A[index] = last + i
			d := dfs(index, last+i)
			if max < d {
				max = d
			}
		}
		return max
	}
	fmt.Println(dfs(0, 0))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// func dfs(A []int)  {

// }
