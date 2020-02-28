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
	m := make(map[int]int)
	keys := make([]int, 0, N)
	max := -1000000009
	for i := 0; i < N; i++ {
		x := nextInt()
		m[x] = nextInt()
	}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, x := range keys {
		// if x-m[x] >= max {
		// 	max = x+m[x]
		// } else {
		// 	delete(m, x)
		// }
		if max > x-m[x] {
			delete(m, x)
		} else {
			max = x + m[x]
		}
	}
	// for i := 0; i < N; i++ {
	// 	// if keys[i]-m[keys[i]] >= max {
	// 	// 	max = keys[i-1] + m[keys[i-1]]
	// 	// }
	// 	// if max > keys[i]-m[keys[i]] {
	// 	// 	delete(m, keys[i])
	// 	// } else {
	// 	// 	max = keys[i]+m[keys[i]]
	// 	// }
	// 	if keys[i]-m[keys[i]] >= max {
	// 		max = keys[i]+m[keys[i]]
	// 	} else {
	// 		delete(m, keys[i])
	// 	}
	// }
	fmt.Println(m)
	fmt.Println(len(m))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
