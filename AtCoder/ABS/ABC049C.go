package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	// dreamerase dreameraser のケースが困る
	// その他考えうる並び
	// eraseerase,eraseeraser,erasedream,erasedreamer
	// erasererase,erasereraser,eraserdream,eraserdreamer
	// dreamdream,dreamdreamer
	// dreamerdream,dreamerdreamer,dreamererase,dreamereraser
	// dreameraserを先にチェック、その次にdreameraseをチェック
	// 消すのではなく、一旦関係ない文字に置換して別の文字列を形成しないようにする
	s = strings.Replace(s, "dreameraser", "X", -1)
	s = strings.Replace(s, "dreamerase", "X", -1)
	s = strings.Replace(s, "dreamer", "X", -1)
	s = strings.Replace(s, "eraser", "X", -1)
	s = strings.Replace(s, "dream", "X", -1)
	s = strings.Replace(s, "erase", "X", -1)
	s = strings.Replace(s, "X", "", -1)
	// fmt.Println(s)
	switch s {
	case "":
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
