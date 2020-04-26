package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

// TODO #10 深さ優先探索の実装もやってみる
func main() {
	N := nextInt()
	P := make([]int, N)
	Q := make([]int, N)
	var a, b int
	for i := range P {
		P[i] = nextInt()
	}
	for i := range Q {
		Q[i] = nextInt()
	}
	Pc := make([]int, N)
	copy(Pc, P)
	sort.Ints(Pc)
	for i := 0; NextPermutation(sort.IntSlice(Pc)); i++ {
		if reflect.DeepEqual(P, Pc) {
			a = i + 1
		}
		if reflect.DeepEqual(Q, Pc) {
			b = i + 1
		}
		if a != 0 && b != 0 {
			break
		}
	}
	fmt.Println(abs(a, b))
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

func fuctorial(a int) int {
	if a == 1 || a == 0 {
		return 1
	} else {
		return fuctorial(a-1) * a
	}
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
