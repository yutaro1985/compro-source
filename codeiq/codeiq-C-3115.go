package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan {
		nums := stringToint(strings.Split(sc.Text(), ""))
		sort.Sort(sort.IntSlice(nums))
		if nums[0] != 0 {
			fmt.Println(nums[3])
		} else {
			fmt.Println(nums[1])
		}
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
