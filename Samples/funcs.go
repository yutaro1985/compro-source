package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000009
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	nums := nextInt()
	trueTotal := 0
	min := 1000000
	var cur int
	for i := 0; i < nums; i++ {
		cur = nextInt()
		if min >= cur {
			trueTotal++
			min = cur
		}
	}
	fmt.Println(trueTotal)
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

func readInt64() int64 {
	sc.Scan()
	r, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return r
}

func readFloat64() float64 {
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), 64)
	return r
}

func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}

// 各桁合計その1 文字で読み込んだとき
func sumNumber(s string) int {
	numbers := strings.Split(s, "")
	var sum int
	for _, sv := range numbers {
		iv, _ := strconv.Atoi(sv)
		sum += iv
	}
	return sum
}

// 各桁合計その2 intで読み込んだとき
func degitsum(n int) int {
	var total int
	for {
		total += n % 10
		n /= 10
		if n == 0 {
			break
		}
	}
	return total
}
