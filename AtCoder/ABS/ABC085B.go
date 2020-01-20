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
	nums := nextInt()
	// appendは遅いらしいので予め必要な容量を確保する
	dn := make([]int, nums)
	for i := 0; i < nums; i++ {
		dn[i] = nextInt()
	}
	dnuniq := intUniq(dn)
	fmt.Println(len(dnuniq))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func intUniq(list []int) []int {
	m := make(map[int]struct{})
	newList := make([]int, 0)

	for _, element := range list {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; ok == false {
			m[element] = struct{}{}
			newList = append(newList, element)
		}
	}
	return newList
}
