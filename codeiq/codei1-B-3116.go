package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	executed := true
	for sc.Scan() {
		nums := stringToint(strings.Split(sc.Text(), ""))
		for {
			nums, executed = DelFirstPair(nums)
			if executed != true {
				break
			}
		}
		for i := range nums {
			fmt.Printf("%v", nums[i])
			if i == len(nums)-1 {
				fmt.Println("")
			}
		}
	}
}

func DelFirstPair(n []int) ([]int, bool) {
	executed := false
	s := make([]int, 0)
	for i := 1; i < len(n); i++ {
		if math.Abs(float64(n[i]-n[i-1])) == 1.0 {
			if i == 1 {
				s = append(s, n[2:]...)
				executed = true
			} else if i == len(n)-1 {
				s = append(s, n[:i-1]...)
				executed = true
			} else {
				s = append(s, n[:i-1]...)
				s = append(s, n[i+1:]...)
				executed = true
			}
			break
		}
	}
	switch executed {
	case true:
		return s, executed
	default:
		return n, executed
	}
}

func Scanner(sc *bufio.Scanner) []string {
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
