package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var odd, even int
	sc := bufio.NewScanner(os.Stdin)
	s := Scanner(sc)
	inputNum := stringToint(s)
	for i := 0; i < inputNum[0]; i++ {
		odd = 0
		even = 0
		s := Scanner(sc)
		cardnums := []rune(s[0])
		for j := len(cardnums) - 2; j >= 0; j -= 2 {
			num, _ := strconv.Atoi(string(cardnums[j]))
			num *= 2
			if num >= 10 {
				num = num/10 + (num % 10)
			}
			even += num
		}
		for k := len(cardnums) - 3; k >= 0; k -= 2 {
			num, _ := strconv.Atoi(string(cardnums[k]))
			odd += num
		}
		x := (10 - (odd+even)%10) % 10
		fmt.Println(x)
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
