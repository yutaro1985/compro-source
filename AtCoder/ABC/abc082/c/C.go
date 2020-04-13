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
	N := nextInt()
	An := make([]int, N)
	var maxAn, ans int
	for i := range An {
		An[i] = nextInt()
		if maxAn < An[i] {
			maxAn = An[i]
		}
	}
	// 普通の配列だといらないところを数えてしまうので、mapを使う
	AnCnt := make(map[int]int)
	for _, Ai := range An {
		AnCnt[Ai]++
	}
	for i, cnt := range AnCnt {
		if cnt > i {
			ans += cnt - i
		} else if cnt < i {
			ans += cnt
		}
	}
	fmt.Println(ans)
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
