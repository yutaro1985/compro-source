package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000009
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	N := nextInt()
	M := nextInt()
	maxleft := 1
	minright := N
	// 全てのMに置いて重複している範囲が分かれば良い
	// 値を格納する必要はなく、一つ一つ比較していけばOK
	for i := 0; i < M; i++ {
		L := nextInt()
		R := nextInt()
		// 重なるところがなかった場合はその時点で0確定
		if maxleft > R || minright < L {
			fmt.Println(0)
			return
		}
		maxleft = max(L, maxleft)
		minright = min(R, minright)
	}
	fmt.Println(minright - maxleft + 1)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// 引数のうち最小のもの
func min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

// 引数のうち最大のもの
func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
