package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// こちらの解き方を参考
// https://kusano-prog.hatenablog.com/entry/20120721/1342874116
// swapしていくと解けるのは実際にあみだを見てみると直感的に理解できる
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func main() {
	var N, L int
	fmt.Scan(&N, &L)
	T := make([]int, N)
	for i := 0; i < N; i++ {
		T[i] = i + 1
	}
	for i := 0; i < L; i++ {
		S := strings.Split(readLine(), "")
		for j := 1; j < N; j++ {
			if S[j*2-1] == "-" {
				T[j-1], T[j] = T[j], T[j-1]
			}
		}
	}
	O := strings.Index(readLine(), "o")
	fmt.Println(T[O/2])
}
