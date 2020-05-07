package main

import (
	"bufio"
	"fmt"
	"os"
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

// 素因数分解の練習

func main() {
	N, P := nextInt(), nextInt()
	pf := primeFuctorize(P)
	ans := 1
	// fmt.Println(pf)
	if N == 1 {
		ans = P
	} else {
		for k, v := range pf {
			// fmt.Println(k, v)
			// ans *= int(math.Pow(float64(k), float64(v/N)))
			ans *= pow(k, v/N)
			// for i := 0; i < v/N; i++ {
			// 	ans *= k
			// 	// fmt.Println("ans", ans, v/N)
			// }
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// https://qiita.com/drken/items/a14e9af0ca2d857dad23#4-%E7%B4%A0%E5%9B%A0%E6%95%B0%E5%88%86%E8%A7%A3
// https://atcoder.jp/contests/caddi2018/submissions/3842181

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
