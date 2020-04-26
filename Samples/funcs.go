package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
	nums := nextInt()
	trueTotal := 0
	min := 1000000
	var cur int
	for i := 0; i < nums; i++ {
		cur = nextInt()
		if min >= cur {
			trueTotal++
			min = cur
		}
	}
	fmt.Println(trueTotal)
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

func readInt64() int64 {
	sc.Scan()
	r, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return r
}

func readFloat64() float64 {
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), 64)
	return r
}

func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}

// 各桁合計その1 文字で読み込んだとき
func sumNumber(s string) int {
	numbers := strings.Split(s, "")
	var sum int
	for _, sv := range numbers {
		iv, _ := strconv.Atoi(sv)
		sum += iv
	}
	return sum
}

// 各桁合計その2 intで読み込んだとき
func degitSum(n int) int {
	var total int
	for {
		total += n % 10
		n /= 10
		if n == 0 {
			break
		}
	}
	return total
}

// 文字配列の重複排除
func strUniq(list []string) []string {
	m := make(map[string]struct{})
	newList := make([]string, 0)

	for _, element := range list {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; ok == false {
			m[element] = struct{}{}
			newList = append(newList, element)
		}
	}
	return newList
}

// Int配列の重複排除
func intUniq(list []int) []int {
	m := make(map[int]struct{})
	newList := make([]int, 0)

	for _, element := range list {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; ok == false {
			m[element] = struct{}{}
			newList = append(newList, element)
		}
	}
	return newList
}

// ここから拝借
// https://qiita.com/kojamam/items/dad162a7360408c9332d
// absをintでラップ
func abs(a int) int {
	return int(math.Abs(float64(a)))
}

// powをintでラップ
// 10の階乗についてはmath.Pow10(n int)を使えばいい
func pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

// 引数のうち最小のもの
// Mathパッケージを使ったものは失敗した…
func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

// 引数のうち最大のもの
func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

// 最大公約数、最小公倍数（2つセット）
func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}

// https://qiita.com/drken/items/0c88a37eec520f82b788#%E5%95%8F%E9%A1%8C-2-abc-148-c---snack-300-%E7%82%B9
// ここを参考に、bをかけるのは最後にした
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
