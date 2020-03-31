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
	Wn := make([]string, N)
	Wnsort := make([]string, N)
	for i := range Wn {
		Wn[i] = nextLine()
	}
	copy(Wnsort, Wn)
	// スライスのコピーはポインタまでコピーしてしまうので下記ではダメ
	// https://qiita.com/riotam/items/06850524450d882e2bd5
	// Wnsort := Wn
	sort.Strings(Wnsort)
	for i := 0; i < N-1; i++ {
		if Wn[i][len(Wn[i])-1] != Wn[i+1][0] || Wnsort[i+1] == Wnsort[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
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
