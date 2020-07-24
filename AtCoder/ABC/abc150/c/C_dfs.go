package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

// dfsで参考に
// https://atcoder.jp/contests/abc150/submissions/14621223
// https://www.qoosky.io/techs/09116bea96

var N int
var P, Q, base []int
var used []bool
var cnt, Pn, Qn int

func main() {
	N = nextInt()
	P = make([]int, N)
	Q = make([]int, N)
	for i := range P {
		P[i] = nextInt()
	}
	for i := range Q {
		Q[i] = nextInt()
	}
	A := make([]int, 0)
	base = make([]int, N)
	copy(base, P)
	sort.Ints(base)
	used = make([]bool, N)
	for i := 0; i < N; i++ {
		dfs(A)
		if Pn != 0 && Qn != 0 {
			fmt.Println(abs(Pn - Qn))
			return
		}
	}
}

func dfs(A []int) {
	if len(A) == N {
		cnt++
		if reflect.DeepEqual(A, P) {
			Pn = cnt
		}
		if reflect.DeepEqual(A, Q) {
			Qn = cnt
		}
		return
	}
	for i := 0; i < N; i++ {
		if !used[i] {
			used[i] = true
			A = append(A, base[i])
			dfs(A)
			A = A[:len(A)-1]
			used[i] = false
		}
	}
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
