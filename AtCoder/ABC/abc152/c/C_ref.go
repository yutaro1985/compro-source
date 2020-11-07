package main
 
import (
	"fmt"
)
 
func main() {
	var n int
	fmt.Scan(&n)
	var cnt int
	min := int(2e5 + 1)
	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)
      fmt.Println(p)
		if p <= min {
			cnt++
			min = p
		}
	}
	fmt.Println(cnt)
}