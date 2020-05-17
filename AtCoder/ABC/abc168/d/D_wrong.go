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

// BFSっぽい実装したつもりだけど、誤り。修正は別ファイルで。

func main() {
	N, M := nextInt(), nextInt()
	// rooms := make([][]int, N)
	// rooms := make([]map[int]int, N)
	// rooms := make(map[int][]int)
	rooms := make([]map[int]bool, N)
	for i := 0; i < N; i++ {
		rooms[i] = make(map[int]bool)
	}
	// for i := 0; i < N; i++ {
	// 	rooms[i] = make(map[int]int)
	// }
	Signposts := make([]int, N)
	prev := make([]int, 0)
	cnt := N
	for i := 0; i < M; i++ {
		A, B := nextInt()-1, nextInt()-1
		rooms[A][B] = true
		rooms[B][A] = true
		if A == 0 {
			Signposts[B] = 1
			prev = append(prev, B)
			cnt--
		}
		if B == 0 {
			Signposts[A] = 1
			prev = append(prev, A)
			cnt--
		}
	}
	// fmt.Println("org", Signposts)
	for cnt >= 0 {
		memo := make([]int, 0)
		for _, v := range prev {
			for aisle := range rooms[v] {
				if aisle != 0 && Signposts[aisle] != 1 {
					Signposts[aisle] = v + 1
					memo = append(memo, aisle)
					cnt--
				}
				if cnt == 0 {
					break
				}
				// fmt.Println(Signposts)
			}
			if cnt == 0 {
				break
			}
		}
		prev := make([]int, 0)
		prev = append(prev, memo...)
	}

	// fmt.Println(rooms, prev, cnt)
	// fmt.Println(Signposts)
	ok := true
	for i := 1; i < N; i++ {
		if Signposts[i] == 0 {
			ok = false
		}
	}
	if ok {
		fmt.Println("Yes")
		for i := 1; i < N; i++ {
			fmt.Println(Signposts[i])
		}
	} else {
		fmt.Println("No")
	}

	// // 順番は保証されそう
	// for _, room := range rooms {
	// 	fmt.Println(room)
	// }
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
