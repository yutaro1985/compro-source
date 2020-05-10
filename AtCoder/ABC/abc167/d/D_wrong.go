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

func main() {
	N, K := nextInt(), nextInt()
	An := make([]int, N)
	sortedAn := make([]int, N+1)
	Revisit := make([]int, N)
	var start, last int
	for i := range An {
		An[i] = nextInt()
	}
	sortedAn[0] = An[0] - 1
	for i := 1; i <= N; i++ {
		sortedAn[i] = An[sortedAn[i-1]] - 1
		if sortedAn[i] == 0 {
			Revisit[An[sortedAn[i-1]]-1] = i
			last = i
			break
		} else if Revisit[An[sortedAn[i-1]]-1] == 0 {
			Revisit[An[sortedAn[i-1]]-1] = -i
		} else {
			start = -Revisit[An[sortedAn[i-1]]-1]
			last = i - 1
			Revisit[An[sortedAn[i-1]]-1] += i
			break
		}
		fmt.Println(sortedAn)
	}
	if K-(start+1) <= 0 {
		fmt.Println(sortedAn[K-1] + 1)
		return
	}
	fmt.Println(K, start, last)
	K -= start
	K %= (last - start + 1)
	fmt.Println(An, sortedAn, K, Revisit)
	sortedAn = sortedAn[start : last+1]
	fmt.Println(sortedAn)
	if K == 0 {
		fmt.Println(1)
		return
	}
	fmt.Println(sortedAn[K-1] + 1)
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
