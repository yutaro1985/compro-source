package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	s := make([]string, N)
	si := make([]int, 1)
	j := 0
	max := 1
	si[0] = 1
	for i := 0; i < N; i++ {
		s[i] = nextLine()
	}
	sort.Strings(s)
	for i := 1; i < N; i++ {
		if s[i] == s[i-1] {
			si[j]++
		} else if j == 0 {
			j++
			si = append(si, 1)
		} else {
			j++
			si = append(si, 1)
		}
		if max < si[j] {
			max = si[j]
		}
	}

	m := 0
	for i := 0; i < len(si); i++ {
		if len(si) == 1 {
			fmt.Println(s[0])
			return
		}
		if si[i] == max {
			fmt.Println(s[m])
		}
		m += si[i]
	}
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
