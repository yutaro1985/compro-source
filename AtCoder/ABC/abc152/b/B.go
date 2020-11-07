package main

import "fmt"
import "strconv"

func main() {
	var a,b string
	var axb,bxa string
	var inta,intb int
	var intaxb,intbxa int
	fmt.Scan(&a)
	fmt.Scan(&b)
	inta, _ = strconv.Atoi(a)
	intb, _ = strconv.Atoi(b)
	for i := 0; i < intb; i++ {
		axb += a
	}
	for i := 0; i < inta; i++ {
		bxa += b
	}
	intaxb, _ = strconv.Atoi(axb)
	intbxa, _ = strconv.Atoi(bxa)
	if intaxb > intbxa {
		fmt.Println(axb)	
	} else {
		fmt.Println(bxa)
	}
}
