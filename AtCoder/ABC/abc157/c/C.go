package main

import (
	"bufio"
	"fmt"
	"math"
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
	var ans int
	N := nextInt()
	M := nextInt()
	conditions := make([]int, N)
	nums := make([]int, N)
	if M == 0 {
		if N == 1 {
			fmt.Println(0)
			return
		}
		nums[0] = 1
		for i := 0; i < N; i++ {
			ans += nums[len(nums)-i-1] * int(math.Pow10(i))
		}
		fmt.Println(ans)
		return
	}
	for i := 0; i < M; i++ {
		s := nextInt() - 1
		c := nextInt()
		if conditions[s] == 1 && nums[s] != c {
			fmt.Println(-1)
			return
		}
		if N != 1 && s == 0 && c == 0 {
			fmt.Println(-1)
			return
		}
		nums[s] = c
		conditions[s] = 1
	}
	if conditions[0] == 0 {
		nums[0] = 1
	}
	if N != 1 && nums[0] == 0 {
		fmt.Println(-1)
		return
	}
	for i := 0; i < N; i++ {
		ans += nums[len(nums)-i-1] * int(math.Pow10(i))
	}
	if ans >= 0 {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
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
