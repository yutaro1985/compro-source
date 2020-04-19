package main

import (
	"bufio"
	"fmt"
	"os"
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
	s := readLine()
	ans := make([]rune, len(s))
	for i, S := range s {
		if 'a' <= S && S <= 'z' {
			ans[i] = S - ('a' - 'A')
		} else if 'A' <= S && S <= 'Z' {
			ans[i] = S + ('a' - 'A')
		} else {
			ans[i] = S
		}
	}
	fmt.Println(string(ans))
}
