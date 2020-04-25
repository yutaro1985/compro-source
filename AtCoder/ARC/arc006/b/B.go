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
	O := strings.Index(readLine(), "o")
	for i := L - 1; i >= 0; i-- {
		if O > 0 && ll[i][O-1] == "-" {
			O -= 2
		} else if O < N*2-2 && ll[i][O+1] == "-" {
			O += 2
		}
	}
	fmt.Println(O/2 + 1)
}
