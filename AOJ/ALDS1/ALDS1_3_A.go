package main

import (
	"bufio"
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

// Popされる順番に注意。
// 減算、除算のときは後からPopした値を前にしないとおかしくなる

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	formula := strings.Split(sc.Text(), " ")
	stack := NewStack()
	for _, v := range formula {
		switch v {
		case "+":
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			stack.Push(b + a)
		case "-":
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			stack.Push(b - a)
		case "*":
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			stack.Push(b * a)
		default:
			num, _ := strconv.Atoi(v)
			stack.Push(num)
		}
	}
	ans, _ := stack.Pop()
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
