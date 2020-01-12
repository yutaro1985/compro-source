//!go
//  tested with : go1.7.4

package main

import "fmt"

func main() {
	var inputNum int
	var inputStrings string
	fmt.Scan(&inputNum)
	for i := 0; i < inputNum; i++ {
		fmt.Scan(&inputStrings)
		fmt.Printf("Hello!, %v\n", inputStrings)
	}
}
