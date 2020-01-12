package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	u := make([]string, 1)
	url_ok := make([]string, 1)
	url_ng := make([]string, 1)

	for sc.Scan() {
		u = append(u, sc.Text())
	}

	for i := 0; i < len(u); i++ {
		_, err := url.Parse(u[i])
		if err != nil {
			url_ng = append(url_ng, u[i])
		} else {
			url_ok = append(url_ok, u[i])
		}
	}

	fmt.Println("- url.Parse OK -")
	for _, url := range URL_OK {
		fmt.Println(url)
	}
	fmt.Println("- url.Parse Err -")
	for _, url := range URL_NG {
		fmt.Println(url)
	}
}
