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
	var check bool

	for i := 1; i < len(inputNum); i++ {

		for j := i + 1; j+1 < inputNum[0]; j++ {
			if inputNum[i]+inputNum[j] == 256 {
				check = true
			}
		}
	}
	if check {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
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
