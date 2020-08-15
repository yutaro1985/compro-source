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

func main() {
	S := nextLine()
	if S == "RRR" {
		fmt.Println(3)
	} else if S == "SSS" {
		fmt.Println(0)
	} else if strings.Contains(S, "RR") {
		fmt.Println(2)
	} else {
		fmt.Println(1)
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

// ここから拝借
// https://shoman.hatenablog.com/entry/2020/02/25/185456
// Stackは[]intのエイリアス
type Stack []int

// Push adds an element
func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

// Pop removes the top element and return it
func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

// Peek returns the top value
func (s *Stack) Peek() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

// Size returns the length of stack
func (s *Stack) Size() int {
	return len(*s)
}

// Empty returns true when stack is empty
func (s *Stack) Empty() bool {
	return s.Size() == 0
}

// NewStack generates stack
func NewStack() *Stack {
	s := new(Stack)
	return s
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
