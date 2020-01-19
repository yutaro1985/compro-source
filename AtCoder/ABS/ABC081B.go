package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	s1 := readLine()
	s2 := readLine()
	s2col := strings.Split(s2, " ")
	nums, _ := strconv.Atoi(s1)
	inputNums := stringToint(s2col)
	times := 0
	flag := false
	for {
		inputNums, flag = div2(nums, inputNums)
		if flag {
			break
		}
		times++
	}
	fmt.Println(times)
}

func div2(len int, nums []int) ([]int, bool) {
	flag := false
	for i, num := range nums {
		if num%2 == 1 {
			flag = true
			break
		} else {
			nums[i] = num / 2
		}
	}
	return nums, flag
}

func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}
