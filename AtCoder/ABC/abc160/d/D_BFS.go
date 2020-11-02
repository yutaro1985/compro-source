package main

import (
	"bufio"
	"container/list"
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

// 参考 ↓の①(BFS)
// https://drken1215.hatenablog.com/entry/2020/04/22/002100

func main() {
	N := nextInt()
	X := nextInt() - 1
	Y := nextInt() - 1
	ans := make([]int, N)
	G := make([][]int, N)
	for i := 0; i < N-1; i++ {
		G[i] = append(G[i], i+1)
		G[i+1] = append(G[i+1], i)
	}
	G[X] = append(G[X], Y)
	G[Y] = append(G[Y], X)
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = -1
		}
	}
	for i := 0; i < N; i++ {
		queue := list.New()
		queue.PushBack(i)
		dist[i][i] = 0
		for queue.Len() != 0 {
			p := queue.Front().Value.(int)
			queue.Remove(queue.Front())
			nps := make([]int, 0)
			if p > 0 {
				nps = append(nps, p-1)
			}
			if p < N-1 {
				nps = append(nps, p+1)
			}
			if p == X {
				nps = append(nps, Y)
			}
			if p == Y {
				nps = append(nps, X)
			}
			for _, np := range nps {
				if dist[i][np] == -1 {
					dist[i][np] = dist[i][p] + 1
					queue.PushBack(np)
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			ans[dist[i][j]]++
		}
	}
	for i := 1; i < N; i++ {
		fmt.Println(ans[i])
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

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
