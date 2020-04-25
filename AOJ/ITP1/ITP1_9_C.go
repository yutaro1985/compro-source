package main

import "fmt"

func main() {
	var n, t, h int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var ts, hs string
		fmt.Scan(&ts, &hs)
		if ts > hs {
			t += 3
		} else if ts < hs {
			h += 3
		} else {
			t++
			h++
		}
	}
	fmt.Println(t, h)
}
