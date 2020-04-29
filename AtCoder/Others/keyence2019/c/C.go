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

// 見当つかなかったので解説見たりググったりした

func main() {
	N := nextInt()
	An := make([]int, N)
	Bn := make([]int, N)
	var sumA, sumB, sumBAdiff, ans int
	ABdiff := make([]int, N)
	for i := range An {
		An[i] = nextInt()
		sumA += An[i]
	}
	for i := range Bn {
		Bn[i] = nextInt()
		sumB += Bn[i]
	}
	if sumA < sumB {
		fmt.Println(-1)
		return
	}
	for i := 0; i < N; i++ {
		if An[i] >= Bn[i] {
			ABdiff[i] = An[i] - Bn[i]
		} else {
			sumBAdiff += Bn[i] - An[i]
			ans++
		}
	}
	if ans == 0 {
		fmt.Println(0)
		return
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ABdiff)))
	for i, d := range ABdiff {
		sumBAdiff -= d
		if sumBAdiff <= 0 {
			fmt.Println(ans + i + 1)
			return
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
