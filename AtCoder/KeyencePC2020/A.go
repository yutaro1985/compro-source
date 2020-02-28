package main

import "fmt"

func main()  {
	var H,W,N,max int
	fmt.Scan(&H,&W,&N)
	if H > W {
		max = H
	} else {
		max = W
	}
	if N%max != 0 {
		fmt.Println(N/max+1)
	} else {
		fmt.Println(N/max)
	}

}
