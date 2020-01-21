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

// i+1番目のtとi番目のtの差は、次の目的地までの移動可能回数と読み替えることができる
// 次の目的地までのx,yの差を取って絶対値を取り、距離を出す。
// 最短距離でそこに行き、一度到達できたらそこから偶数回移動できればその場所に行ける
// 残り移動回数が奇数回だとたどり着けない。
// 下記を満たせば到達可能
// t > x + y かつ t - (x + y ) ≡ 0 mod 2

func main() {
	nums := nextInt()
	// 初期値は原点
	prev := []int{0, 0, 0}
	for i := 0; i < nums; i++ {
		nextgoal := readNums()
		// 次の目的地までの所要時間と距離
		nextdist := []int{0, 0, 0}
		// 時間は必ず次の値が大きくなるので絶対値とらない
		nextdist[0] = nextgoal[0] - prev[0]
		x := nextgoal[1] - prev[1]
		y := nextgoal[2] - prev[2]
		nextdist[1] = int(math.Abs(float64(x)))
		nextdist[2] = int(math.Abs(float64(y)))
		// 次の目的地までの距離を残りの移動可能回数から引いた値が偶数ならたどり着ける
		// また、距離が移動可能回数を
		if (nextdist[0]-(nextdist[1]+nextdist[2]))%2 != 0 || nextgoal[0] < nextgoal[1]+nextgoal[2] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func readNums() []int {
	t := nextInt()
	x := nextInt()
	y := nextInt()
	nums := []int{t, x, y}
	return nums
}
