package main

import "fmt"

// 素数判定
// https://qiita.com/TD12734/items/58551cfab11884e68411

func main() {
	var N, sum int
	var Prime bool
	fmt.Scan(&N)

	for i := N; i != 0; i /= 10 {
		sum += i % 10
	}
	Prime = isPrime(N)
	if !Prime && N%10 != 5 && (N%10)%2 != 0 && sum%3 != 0 {
		Prime = true
	}
	if Prime && N != 1 {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	} else if n%2 == 0 || n < 2 {
		return false
	} else {
		for i := 3; i*i <= n; i += 2 {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}
