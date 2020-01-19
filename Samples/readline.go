package main

import (
	"fmt"
	"bufio"
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
	s1 := readLine()
	s2 := readLine()
	fmt.Println(s1)
	fmt.Println(s2)
}