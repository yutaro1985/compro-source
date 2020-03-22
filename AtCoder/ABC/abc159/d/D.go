package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	N := nextInt()
	An := make([]int, N)
	nums := make([]int, N)
	combi := make([]int, N)
	combi2 := make([]int, N)
	ans := 0
	for i := 0; i < N; i++ {
		An[i] = nextInt()
		nums[An[i]-1]++
	}
	for i := 0; i < N; i++ {
		combi[i] = combination2(nums[i])
		if nums[i] != 0 {
			combi2[i] = combination2(nums[i] - 1)
		} else {
			combi2[i] = 0
		}
	}
	for i := 0; i < N; i++ {
		ans += combi[i]
	}
	for i := 0; i < N; i++ {
		fmt.Println(ans - (combi[An[i]-1] - combi2[An[i]-1]))
	}
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func combination2(num int) int {
	return num * (num - 1) / 2
}
