package main

import "fmt"

func main() {
    fmt.Println(succ("9"))    // => "10"
    fmt.Println(succ("a"))    // => "b"
    fmt.Println(succ("AAA"))  // => "AAB"
    fmt.Println(succ("A99"))  // => "B00"
    fmt.Println(succ("A099")) // => "A100"
	// fmt.Println(reverse("ABC"))
}

// 文字列を反転して返す
func reverse(s string) string {
    ans := ""
    for i, _ := range s {
        ans += string(s[len(s)-i-1])
    }
    return string(ans)
}

// "次"の文字列を取得する
func succ(s string) string {
    r := reverse(s)
    ans := ""
    carry := 1
    lastLetter := string(r[0])
    for i, _ := range r {
        lastLetter = string(r[i])
        a := lastLetter
        if carry == 1 {
            if lastLetter == "z" {
                a = "a"
                carry = 1
            } else if lastLetter == "Z" {
                a = "A"
                carry = 1
            } else if lastLetter == "9" {
                a = "0"
                carry = 1
            } else {
                if r[i] == 0 {
                    a = "1"
                } else {
                    a = string(r[i] + 1)
                    carry = 0
                }
            }
        }
        ans += a
    }
    if carry == 1 {
        if lastLetter == "9" {
            ans += "1"
        } else if lastLetter == "z" {
            ans += "a"
        } else if lastLetter == "Z" {
            ans += "A"
        }
    }
    return reverse(ans)
}