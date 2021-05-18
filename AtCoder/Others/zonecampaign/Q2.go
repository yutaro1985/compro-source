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

var buf []byte = make([]byte, initialBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

// 問題によって値は調整する
const (
	mod     = int(1e9) + 7
	maxsize = 510000
	INF     = 1 << 60
)

func main() {
	var ans int
	Sieve := MakeSieve(1000)
	for i := 0; i < 20*20; i++ {
		m := nextInt()
		if m >= 2 && Sieve[m] == m {
			ans++
		}
	}
	fmt.Println(ans)
}

// FastPrimeFactorize osa_k法を使った高速素因数分解
// sieveを作るところのjはi*2スタートでもよい？
// https://qiita.com/rsk0315_h4x/items/ff3b542a4468679fb409
func FastPrimeFactorize(n int, sieve []int) map[int]int {
	pf := make(map[int]int)
	for n > 1 {
		pf[sieve[n]]++
		n /= sieve[n]
	}
	return pf
}

// MakeSieve は高速素数判定に使う篩を作る。
// Sieve[i] < i なら合成数なので、Sieveは素数判定にも使える
func MakeSieve(max int) []int {
	sieve := make([]int, max)
	for i := range sieve {
		sieve[i] = i
	}
	for i := 2; i*i <= max; i++ {
		if sieve[i] == i {
			for j := i; j < max; j += i {
				if sieve[j] == j {
					sieve[j] = i
				}
			}
		}
	}
	return sieve
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
