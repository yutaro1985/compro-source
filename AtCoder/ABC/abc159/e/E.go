package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

// ビット演算のサンプルとして丁度いい問題？
// TODO #6 写経してみたがあまり理屈がわかっておらず、微妙に正解にもならないので後で見直す

func main() {
	H := nextInt()
	W := nextInt()
	K := nextInt()
	s := make([][]int, H)
	ans := int(1e10)
	// c := make([][]int, H)
	// for i := 0; i < H; i++ {
	// 	c[i] = make([]int, W)
	// }
	for i := 0; i < H; i++ {
		s[i] = make([]int, W)
		nums := strings.Split(nextLine(), "")
		for j := 0; j < W; j++ {
			s[i][j], _ = strconv.Atoi(nums[j])
		}
	}
	// シフト演算で2のn乗
	// 1<<uint(H-1)
	for div := 0; div < int(math.Pow(2.0, float64(H-1))); div++ {
		g := 0
		id := make([]int, H)
		for j := 0; j < H; j++ {
			if div>>uint(j) == 1 {
				g++
			}
		}
		g++
		c := make([][]int, H)
		for i := 0; i < H; i++ {
			c[i] = make([]int, W)
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				c[id[i]][j] += s[i][j]
			}
		}
		num := g - 1
		cur := make([]int, g)
		for j := 0; j < W; j++ {
			// for i := 0; i < g; i++ {
			// 	cur[i] += c[i][j]
			// }
			// check := true
			// for i := 0; i < g; i++ {
			// 	if cur[i] > K {
			// 		check = false
			// 	}
			// 	if !check {
			// 		num++
			// 		cur = make([]int, g)

			// 	}
			// }
			if !add(j, cur, c, g, K) {
				num++
				cur = make([]int, g)
				if !add(j, cur, c, g, K) {
					num = int(1e10)
					break
				}
			}
		}
		ans = MinOf(ans, num)
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

func add(j int, cur []int, c [][]int, g int, K int) bool {
	for i := 0; i < g; i++ {
		cur[i] += c[i][j]
	}
	for i := 0; i < g; i++ {
		if cur[i] > K {
			return false
		}
	}
	return true
}

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}
