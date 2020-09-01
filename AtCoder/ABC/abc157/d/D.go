package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

var d = []Position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var d8 = []Position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func main() {
	N, M, K := nextInt(), nextInt(), nextInt()
	friends := make([]map[int]bool, N)
	for i := range friends {
		friends[i] = make(map[int]bool)
	}
	blocklist := make([][]int, N)
	uf := newUnionFind(N)
	for i := 0; i < M; i++ {
		A, B := nextInt()-1, nextInt()-1
		friends[A][B] = true
		friends[B][A] = true
		uf.unite(A, B)
	}
	for i := 0; i < K; i++ {
		C, D := nextInt()-1, nextInt()-1
		blocklist[C] = append(blocklist[C], D)
		blocklist[D] = append(blocklist[D], C)
	}
	apfriends := make([]int, N)
	for i := 0; i < N; i++ {
		ans := uf.size(i) - 1 - len(friends[i])
		for _, v := range blocklist[i] {
			if uf.same(i, v) {
				ans--
			}
		}
		apfriends[i] = ans
	}

	fmt.Println(strings.Trim(fmt.Sprint(apfriends), "[]"))
}

// Position として迷路問題での現在地を表す構造体を定義
type Position struct {
	H int
	W int
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

func ceil(a, b int) int {
	return (a + (b - 1)) / b
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

// golangの公式サンプルより
// https://xn--go-hh0g6u.com/pkg/container/heap/#example__intHeap
// IntHeap は，整数の最小ヒープです。
type IH []int

func (h IH) Len() int           { return len(h) }
func (h IH) Less(i, j int) bool { return h[i] < h[j] }
func (h IH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IH) Push(x interface{}) {
	// Push と Pop はポインタレシーバを使っています。
	// なぜなら，スライスの中身だけでなく，スライスの長さも変更するからです。
	*h = append(*h, x.(int))
}

func (h *IH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// golangの公式サンプルより
// https://xn--go-hh0g6u.com/pkg/container/heap/#example__priorityQueue

// Item は，優先キューで管理する項目です。
type Item struct {
	value    string // 値。任意です。
	priority int    // キューにおける優先度
	// index は heap.Interface メソッドで更新されます。
	index int // ヒープにおけるインデックス
}

// PriorityQueue は heap.Interface を実装し，Item のリストを保持します。
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Pop が最小ではなく最大の優先度を持つ項目を返して欲しいので，ここでは > を使っています。
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // メモリリークを避ける
	item.index = -1 // 安全のため
	*pq = old[0 : n-1]
	return item
}

// update はキューの Item の優先度と値を更新します。
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// UnionFind の定義
// 下記を参考
// https://youtu.be/TdR816rqc3s?t=6822
// https://github.com/atcoder/live_library/blob/master/uf.cpp
// https://qiita.com/haru1843/items/2295d0ec1f5002bd5c33#%E5%AE%9F%E8%A3%85
type UnionFind struct {
	parent []int
	maxlen int
}

// UnionFind のスライス初期化
func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.parent = make([]int, N)
	u.maxlen = 1
	for i := range u.parent {
		u.parent[i] = -1
	}
	return u
}

// xの根を見つける
func (u *UnionFind) find(x int) int {
	if u.parent[x] < 0 {
		return x
	}
	u.parent[x] = u.find(u.parent[x])
	return u.parent[x]
}

// xとyのグループを結合する
func (u *UnionFind) unite(x, y int) {
	xf := u.find(x)
	yf := u.find(y)
	if xf == yf {
		return
	}
	// 常に大きい方に小さい木をくっつける
	if xf > yf {
		xf, yf = yf, xf
	}
	u.parent[xf] += u.parent[yf]
	u.parent[yf] = xf
	if u.parent[xf] < 0 {
		u.maxlen = MaxOf(u.maxlen, abs(u.parent[xf]))
	}
}

// xとyが同じグループに所属するかどうかを返す
func (u *UnionFind) same(x, y int) bool {
	if u.find(x) == u.find(y) {
		return true
	}
	return false
}

// xの所属するグループの木の大きさを返す
func (u *UnionFind) size(x int) int {
	return -u.parent[u.find(x)]
}

// 根の数を返す
func (u UnionFind) rootcnt() int {
	var cnt int
	for _, v := range u.parent {
		if v < 0 {
			cnt++
		}
	}
	return cnt
}
