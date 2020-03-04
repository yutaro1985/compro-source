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
	s := readLine()
	s = strings.Replace(s, "Left", "<", -1)
	s = strings.Replace(s, "Right", ">", -1)
	s = strings.Replace(s, "AtCoder", "A", -1)
	fmt.Println(s)
}
