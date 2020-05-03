package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	N := nextInt()
	var ans, sum int
	scores := make([]int, N)
	for i := range scores {
		scores[i] = nextInt()
		sum += scores[i]
	}
	sort.Ints(scores)
	if sum%10 != 0 {
		ans = sum
	} else {
		for i := 0; i < N; i++ {
			if scores[i]%10 != 0 {
				ans = sum - scores[i]
				break
			}
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
