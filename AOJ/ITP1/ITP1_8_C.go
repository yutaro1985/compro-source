package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1行ずつEOFまで読み込み、というときの例
	sc := bufio.NewScanner(os.Stdin)
	letters := make([]byte, 'z'-'a'+1)
	for sc.Scan() {
		S := strings.ToLower(sc.Text())
		for _, s := range S {
			if 'a' <= s && s <= 'z' {
				letters[s-'a']++
			}
		}
	}
	for i, num := range letters {
		fmt.Printf("%s : %d\n", string('a'+i), num)
	}
}
