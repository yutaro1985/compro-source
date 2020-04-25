package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	ll := make([][]string, L)
	for i := 0; i < L; i++ {
		ll[i] = make([]string, 2*N)
		ll[i] = strings.Split(readLine(), "")
	}
	O := strings.Split(readLine(), "")
	for i := 0; i < N*2; i += 2 {
		ans := i/2 + 1
		loc := i
		moved := false
		for j := 0; j < L; j++ {
			if loc != N*2-2 {
				if ll[j][loc+1] == "-" {
					loc += 2
					moved = true
				}
			}
			if loc != 0 && !moved {
				if ll[j][loc-1] == "-" {
					loc -= 2
				}
			}
		}
		if O[loc] == "o" {
			fmt.Println(ans)
			return
		}
	}
}
