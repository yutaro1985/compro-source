package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, L int
	var ans string
	fmt.Scan(&N, &L)
	Sn := make([]string, N)
	for i := range Sn {
		fmt.Scan(&Sn[i])
	}
	sort.Strings(Sn)
	for _, s := range Sn {
		ans += s
	}
	fmt.Println(ans)
}
