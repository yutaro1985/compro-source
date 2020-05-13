package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	Deg, Dis := nextInt(), nextInt()
	ws := windSpeed(Dis)
	if ws == 0 {
		fmt.Println("C", 0)
	} else {
		fmt.Println(checkDirection(Deg), windSpeed(Dis))
	}
}

func checkDirection(Deg int) string {
	var direction string
	if 113 <= Deg && Deg < 338 {
		direction = "NNE"
	} else if 338 <= Deg && Deg < 563 {
		direction = "NE"
	} else if 563 <= Deg && Deg < 788 {
		direction = "ENE"
	} else if 788 <= Deg && Deg < 1013 {
		direction = "E"
	} else if 1013 <= Deg && Deg < 1238 {
		direction = "ESE"
	} else if 1238 <= Deg && Deg < 1463 {
		direction = "SE"
	} else if 1463 <= Deg && Deg < 1688 {
		direction = "SSE"
	} else if 1688 <= Deg && Deg < 1913 {
		direction = "S"
	} else if 1913 <= Deg && Deg < 2138 {
		direction = "SSW"
	} else if 2138 <= Deg && Deg < 2363 {
		direction = "SW"
	} else if 2363 <= Deg && Deg < 2588 {
		direction = "WSW"
	} else if 2588 <= Deg && Deg < 2813 {
		direction = "W"
	} else if 2813 <= Deg && Deg < 3038 {
		direction = "WNW"
	} else if 3038 <= Deg && Deg < 3263 {
		direction = "NW"
	} else if 3263 <= Deg && Deg < 3488 {
		direction = "NNW"
	} else {
		direction = "N"
	}
	return direction
}

func windSpeed(Dis int) int {
	var W int
	Disf := math.Round(float64(Dis)/60*10) / 10
	if 0.0 <= Disf && Disf <= 0.2 {
		W = 0
	} else if 0.3 <= Disf && Disf <= 1.5 {
		W = 1
	} else if 1.6 <= Disf && Disf <= 3.3 {
		W = 2
	} else if 3.4 <= Disf && Disf <= 5.4 {
		W = 3
	} else if 5.5 <= Disf && Disf <= 7.9 {
		W = 4
	} else if 8.0 <= Disf && Disf <= 10.7 {
		W = 5
	} else if 10.8 <= Disf && Disf <= 13.8 {
		W = 6
	} else if 13.9 <= Disf && Disf <= 17.1 {
		W = 7
	} else if 17.2 <= Disf && Disf <= 20.7 {
		W = 8
	} else if 20.8 <= Disf && Disf <= 24.4 {
		W = 9
	} else if 24.5 <= Disf && Disf <= 28.4 {
		W = 10
	} else if 28.5 <= Disf && Disf <= 32.6 {
		W = 11
	} else {
		W = 12
	}
	return W
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func pow(p, q int) int {
	res := p
	if q == 0 {
		return 1
	}
	for i := 0; i < q-1; i++ {
		res *= p
	}
	return res
}

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}

func lcmof2numbers(a int, b int) int {
	return a / gcdof2numbers(a, b) * b
}

// マイナスの場合は考慮していない
func fuctorial(a int) int {
	if a == 1 || a == 0 {
		return 1
	} else {
		return fuctorial(a-1) * a
	}
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

// 素因数分解したmapを返す
func primeFuctorize(n int) map[int]int {
	pf := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			pf[i]++
			n /= i
		}
	}
	if n != 1 {
		pf[n]++
	}
	return pf
}
