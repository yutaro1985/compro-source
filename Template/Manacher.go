// Manacher はManacher Algorithmを用いて文字列sに含まれる最長の回文を返す
func Manacher(s string) string {
	preProcess := func(s string) string {
		n := len(s)
		if n == 0 {
			return "^$"
		}
		result := "^"
		for i := 0; i < n; i++ {
			result += "#" + string(s[i])
		}
		result += "#$"
		return result
	}
	t := preProcess(s)
	n := len(t)
	p := make([]int, n)
	c, r := 0, 0

	for i := 1; i < n-1; i++ {
		mirror := 2*c - i
		if i < r {
			p[i] = MinOf(r-i, p[mirror])
		}

		for t[i+(1+p[i])] == t[i-(1+p[i])] {
			p[i]++
		}

		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i := 1; i < n-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	start := (centerIndex - 1 - maxLen) / 2
	end := start + maxLen
	return s[start:end]
}
