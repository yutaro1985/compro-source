package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// このソース自体は多分動かない。

const (
	// 初期バッファサイズ
	initialBufSize = 10000
	// バッファサイズの最大値。Scannerは必要に応じこのサイズまでバッファを大きくして各行をスキャンする。
	// この値がinitialBufSize以下の場合、Scannerはバッファの拡張を一切行わず与えられた初期バッファのみを使う。
	maxBufSize = 1000000
)

var sc = bufio.NewScanner(os.Stdin)

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

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func main() {
	var s1, s2 string
	s1 = nextLine()
	s2 = nextLine()
	nums, _ := strconv.Atoi(s1)
	slice := strings.Split(s2, " ")
	inputNum := stringToint(slice)

	truenums := 0
	min := inputNum[0]
	for i := 0; i < nums; i++ {
		if min >= inputNum[i] {
			min = inputNum[i]
			truenums++
		}
	}
	fmt.Println(truenums)
}
func Scanner() []string {
	sc := bufio.NewScanner(os.Stdin)
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	slice := strings.Split(s, " ")
	return slice
}

func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}

// 一つずつ読み込んで比較するのでもしかすると時間かかるかも
func max(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func min(a []int) int {
	min := a[0]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}
