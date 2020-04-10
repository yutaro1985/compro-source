package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	N := nextInt()
	An := make([]int, N)
	var sum, cost1, cost2, ans int
	var ave float64
	for i := range An {
		An[i] = nextInt()
		sum += An[i]
	}
	ave = float64(sum) / float64(N)
	for _, Ai := range An {
		ave1 := int(math.Ceil(ave))
		ave2 := int(math.Floor(ave))
		cost1 += cost(Ai, ave1)
		cost2 += cost(Ai, ave2)
	}
	if cost1 < cost2 {
		ans = cost1
	} else {
		ans = cost2
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

func cost(Ai int, ave int) int {
	if Ai > ave {
		return (Ai - ave) * (Ai - ave)
	} else {
		return (ave - Ai) * (ave - Ai)
	}
}
