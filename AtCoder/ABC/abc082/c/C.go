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
	ans := 0
	AnCnt := make(map[int]int)
	// 普通の配列だといらないところを数えてしまうので、mapを使う
	for i := 0; i < N; i++ {
		AnCnt[nextInt()]++
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
