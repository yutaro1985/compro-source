package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	_, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	an := strings.Split(sc.Text(), " ")
	ans := ""
	for i := len(an) - 1; i >= 0; i-- {
		ans += an[i] + " "
	}
	fmt.Println(strings.TrimSpace(ans))
}
