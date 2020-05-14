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

type person struct {
	id int
	x  []int
	y  []int
}

// bit全探索の練習
// https://atcoder.jp/contests/abc147/tasks/abc147_c

func main() {
	N := nextInt()
	people := make([]person, N)
	var ans int
	for i := 0; i < N; i++ {
		A := nextInt()
		for j := 0; j < A; j++ {
			people[i].x = append(people[i].x, nextInt()-1)
			people[i].y = append(people[i].y, nextInt())
		}
	}
	for bit := 0; bit < 1<<uint(N); bit++ {
		var cnt int
		ok := true
		for i := 0; i < N; i++ {
			// i人目が1であることを前提に調べているので、
			// i人目が0の場合は見なくて良い
			if uint(bit)>>uint(i)&1 == 1 {
				for j := 0; j < len(people[i].x); j++ {
					// xorを使って一つにまとめられる
					if uint(people[i].y[j])^uint(bit)>>uint(people[i].x[j])&1 == 1 {
						ok = false
						break
					}
					// if people[i].y[j] == 1 && uint(bit)>>uint(people[i].x[j])&1 == 0 {
					// 	ok = false
					// 	break
					// }
					// if people[i].y[j] == 0 && uint(bit)>>uint(people[i].x[j])&1 == 1 {
					// 	ok = false
					// 	break
					// }
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			for i := 0; i < N; i++ {
				if uint(bit)>>uint(i)&1 == 1 {
					cnt++
				}
			}
			ans = MaxOf(ans, cnt)
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
