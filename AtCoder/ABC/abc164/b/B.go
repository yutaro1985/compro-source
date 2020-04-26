package main

import "fmt"

func main()  {
	var A,B,C,D int
	fmt.Scan(&A,&B,&C,&D)
	for {
		C -= B
		if C <= 0 {
			fmt.Println("Yes")
			return
		}
		A -= D
		if A <= 0 {
			fmt.Println("No")
			return
		}
	}
}