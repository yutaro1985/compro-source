package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	sl := strings.Split(s, "")
	tl := strings.Split(t, "")
	sort.Strings(sl)
	sort.Sort(sort.Reverse(sort.StringSlice(tl)))
	s = strings.Join(sl, "")
	t = strings.Join(tl, "")
	if strings.Compare(s, t) < 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
