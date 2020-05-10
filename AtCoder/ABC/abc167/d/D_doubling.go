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
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

// 解説放送より
// https://youtu.be/ENSOy8u9K9I?t=3667

const (
	D     int = 60
	MAX_N     = 200005
)

func main() {
	N, K := nextInt(), nextInt()
	// An := make([]int, N)
	to := make([][]int, D)
	for i := range to {
		to[i] = make([]int, MAX_N)
	}
	// to は jから2^i個先
	for i := 0; i < N; i++ {
		// ここで-1したので最後に+1する
		to[0][i] = nextInt() - 1
	}
	for i := 0; i < D-1; i++ {
		for j := 0; j < N; j++ {
			// ここの計算は表を見るとわかりやすい
			// 前の行(i行目）の、左からj番目の数字を見て、i行目の左から数えてその数字の位置にある数字を次の行に入れる
			to[i+1][j] = to[i][to[i][j]]
		}
	}
	cur := 0
	for i := D - 1; i >= 0; i-- {
		// 2^iをシフト演算で。
		l := 1 << uint(i)
		if l <= K {
			cur = to[i][cur]
			K -= l
		}
	}
	fmt.Println(cur + 1)
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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func pow(p, q int) int {
	res := p
	if q == 0 {
		return 1
	}
	for i := 0; i < q-1; i++ {
		res *= p
	}
	return res
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

func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}

func lcmof2numbers(a int, b int) int {
	return a / gcdof2numbers(a, b) * b
}

// マイナスの場合は考慮していない
func fuctorial(a int) int {
	if a == 1 || a == 0 {
		return 1
	} else {
		return fuctorial(a-1) * a
	}
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

// 素因数分解したmapを返す
func primeFuctorize(n int) map[int]int {
	pf := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			pf[i]++
			n /= i
		}
	}
	if n != 1 {
		pf[n]++
	}
	return pf
}
