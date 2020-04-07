package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e8 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	N := nextInt()
	K := nextInt()
	An := make([]int, N)
	nums := []int{0}
	var count, prev, ans int
	// まずすべての数字を読み取って小さい順ソート
	for i := range An {
		An[i] = nextInt()
	}
	sort.Ints(An)
	// 小さい順にソートしてから、一つずつカウント
	// 配列の長さが数字の種類
	for i := 0; i < N; i++ {
		if An[i] != prev {
			prev = An[i]

			nums = append(nums, 0)
			count++
		}
		nums[count]++
	}
	// 配列の2番目から数字が入るので、2番目以降を切り出したスライスを再定義
	numsvars := nums[1:]
	// 数字の数が少ない順にソート
	sort.Ints(numsvars)
	// 数字の種類＝ソートした配列の長さ-K番目までの数字を足した分が消す必要のある数字の数
	if K < len(numsvars) {
		for i := 0; i < len(numsvars)-K; i++ {
			ans += numsvars[i]
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
