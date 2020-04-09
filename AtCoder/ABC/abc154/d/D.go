package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e8 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

// 期待値・累積和
// https://drken1215.hatenablog.com/entry/2020/02/10/235800
// https://qiita.com/drken/items/56a6b68edef8fc605821

func main() {
	N := nextInt()
	K := nextInt()
	Pn := make([]int, N)
	Csum := make([]int, N+1)
	var ans int
	// (p+1)/2 がサイコロの期待値。2で割るのは最後にして、まずはpi+1を配列に入れる
	for i := range Pn {
		Pn[i] = nextInt()
		Pn[i]++
	}
	// 累積和
	for i, Pi := range Pn {
		Csum[i+1] = Csum[i] + Pi
	}

	for i := 0; i+K <= N; i++ {
		sum := Csum[i+K] - Csum[i]
		if ans < sum {
			ans = sum
		}
	}
	// 以下はどっちでも通った
	// fmt.Printf("%.7f\n", float64(sans)/2)
	fmt.Println(float64(ans) / 2)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
