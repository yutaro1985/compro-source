// ReverseString は文字列を反転させたものを返す
func ReverseString(S string) string {
	res := []rune(S)
	for i, j := 0, len(S)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}