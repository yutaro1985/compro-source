ng,ok := -1,N

func isOK() bool  {
	
}

func bisearch(key int) int {
	for Abs(ok-ng) >1 {
		mid := (ok+ng)/2
		if isOK() {
			ok = mid
		} else {
			ng = mid
		}
	}	
	return ok
}
