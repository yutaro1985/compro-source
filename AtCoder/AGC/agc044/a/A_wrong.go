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

func main() {
	T := nextInt()
	// pfs := map[int]map[int]int{}
	// for i := 0; i <= 60; i++ {
	// 	for j := 0; j <= 38; j++ {
	// 		for k := 0; k <= 26; k++ {
	// 			product := pow(2, i) * pow(3, j) * pow(5, k)
	// 			pfs[product] = map[int]int{
	// 				2: i,
	// 				3: j,
	// 				5: k,
	// 			}
	// 		}
	// 	}
	// }
	for testcases := 0; testcases < T; testcases++ {
		N, A, B, C, D := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
		var x, ans int
		x = N
		for x > 0 {
			for i := 0; i < 60; i++ {
				for j := 0; j < 38; j++ {
					for k := 0; k < 26; k++ {

					}
				}
			}
		}
		// ans := int(1e18)
		// useCoins := map[int]int{
		// 	2: A,
		// 	3: B,
		// 	5: C,
		// }
		// // fmt.Println(useCoins[2], useCoins[3], useCoins[5], ans)
		// for i := N - 100000; i <= N+100000; i++ {
		// 	if i <= 0 {
		// 		continue
		// 	}
		// 	if i%2 == 0 || i%3 == 0 || i%5 == 0 {
		// 		tmp := abs(N-i)*D + D
		// 		if tmp > MinOf(N*D, math.MaxInt64) {
		// 			continue
		// 		}
		// 		pf := pfs[i]
		// 		// fmt.Println(i, tmp, pf)
		// 		ok := true
		// 		if len(pf) == 0 {
		// 			continue
		// 		}
		// 		for k, v := range pf {
		// 			if k > 5 {
		// 				ok = false
		// 				break
		// 			}
		// 			tmp += v * useCoins[k]
		// 		}
		// 		if !ok {
		// 			continue
		// 		}
		// 		ans = MinOf(ans, tmp)
		// 		// fmt.Println(ans, tmp)
		// 	}
		// }
		fmt.Println(ans)
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

func primeFuctorize235(n int) map[int]int {
	pf := map[int]int{}
	for i := 2; i <= 5; i++ {
		for n%i == 0 {
			pf[i]++
			n /= i
		}
	}
	if n != 1 {
		pf[n]++
		if n != 2 && n != 3 && n != 5 {
			pf = map[int]int{}
		}
	}
	return pf
}
