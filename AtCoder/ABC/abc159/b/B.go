package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	ans := "No"
	first := S[0 : (len(S)-1)/2]
	last := S[(len(S)+3)/2-1 : len(S)]
	if paindrome(S) && paindrome(first) && paindrome(last) {
		ans = "Yes"
	}
	fmt.Println(ans)
}

// 回文判定
func paindrome(s string) bool {
	return s == Reverse(s)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
