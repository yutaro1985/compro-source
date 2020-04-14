package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	K := nextInt()
	S := nextInt()
	// stringの+=は遅い
	ans := make([]string, 0)
	for i := 0; i < N; i++ {
		if i < K {
			ans = append(ans, strconv.Itoa(S))
		} else if S != 1e9 {
			ans = append(ans, strconv.Itoa(S+1))
		} else {
			ans = append(ans, "1")
		}
	}
	fmt.Println(strings.Join(ans, " "))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
