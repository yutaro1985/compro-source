package main

import "fmt"
import "strconv"

func main() {
	s := "123"
	// strconvで数字を文字列として扱い、戻すときのやり方
	// https://xn--go-hh0g6u.com/pkg/strconv/
	// Atoiはエラー時の戻り値を処理する必要がある
	s0, _ := strconv.Atoi(string(s[0]))
	s1, _ := strconv.Atoi(string(s[1]))
	s2, _ := strconv.Atoi(string(s[2]))
	fmt.Println(len(s))
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
}
