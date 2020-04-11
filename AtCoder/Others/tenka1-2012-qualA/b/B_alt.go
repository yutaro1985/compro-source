package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1e6)

func main() {
	cn := readLine()
	var ans []rune
	for i, ci := range cn {
		if !unicode.IsSpace(ci) {
			if i != 0 && cn[i-1] == ' ' {
				ans = append(ans, ',')
			}
			ans = append(ans, ci)
		}
	}
	fmt.Println(string(ans))
}

func readLine() string {
	buf := make([]byte, 0, 1e6)
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
