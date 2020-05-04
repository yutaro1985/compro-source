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
// とても参考になる記事
// https://drken1215.hatenablog.com/entry/2020/05/04/190252

type nums struct {
	A int
	B int
	C int
	D int
}

func main() {
	N, M, Q := nextInt(), nextInt(), nextInt()
	nums := make([]nums, Q)
	A := make([]int, 0)
	for i := range nums {
		nums[i].A = nextInt() - 1
		nums[i].B = nextInt() - 1
		nums[i].C = nextInt()
		nums[i].D = nextInt()
	}
	fmt.Println(dfs(A, N, M, nums))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func dfs(A []int, size int, maxnum int, nums []nums) int {
	if len(A) == size {
		return score(A, nums)
	}
	// Ai <= Ai+1なので、一つ前の項の最高値を次の項のスタートにする
	var prevmax, max int
	if len(A) != 0 {
		prevmax = A[len(A)-1]
	}
	for i := prevmax; i < maxnum; i++ {
		A = append(A, i)
		maxscore := dfs(A, size, maxnum, nums)
		if max < maxscore {
			max = maxscore
		}
		A = A[:len(A)-1]
	}
	return max
}

// 生成した配列が条件に沿った長さになったらさせる処理を関数として外に出している
func score(A []int, nums []nums) int {
	var score int
	for i := 0; i < len(nums); i++ {
		if A[nums[i].B]-A[nums[i].A] == nums[i].C {
			score += nums[i].D
		}
	}
	return score
}
