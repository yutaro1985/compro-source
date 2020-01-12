package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := Scanner()
	inputNum := stringToint(s)
	count := 0
	numOfTrue, _ := strconv.ParseUint(s[1],10,64)
	for i := 0; i <= inputNum[0]; i++ {
		bits := fmt.Sprintf("%b", i)
		bitsUint, _ := strconv.ParseUint(bits,2,64)
		bitsUint = numOfBits(bitsUint)
		if bitsUint == numOfTrue {
			count++
		}
	}
	fmt.Println(count)
}

func Scanner() []string {
	sc := bufio.NewScanner(os.Stdin)
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	slice := strings.Split(s, " ")
	return slice
}

func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}

func numOfBits(bits uint64) uint64 {
	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >>16 & 0x0000ffff)
}