package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scan(&s)
	// 正規表現使うとあっという間
	// ただ、正規表現の書き方が合ってるかを検証するのがめんどくさそう…。
	// 前方一致、後方一致を指定しないと失敗する
	var re = regexp.MustCompile(`^(dream(er)?|eraser?)+$`)
	// var re = regexp.MustCompile(`^(dream|dreamer|erase|eraser)+$`)
	if re.MatchString(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
