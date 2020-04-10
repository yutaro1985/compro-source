package main

import "fmt"

func main() {
	var s, ans string
	fmt.Scan(&s)
	for _, letter := range s {
		if letter != 'B' {
			ans += string(letter)
		} else if ans != "" {
			ans = ans[:len(ans)-1]
		}
	}
	fmt.Println(ans)
}
