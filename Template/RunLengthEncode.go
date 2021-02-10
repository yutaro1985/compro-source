type RL struct {
	Str byte
	Num int
}

// RunLengthEncode は文字列を同じ文字がいくつ続けて現れるかごとにPairにしたスライスに格納して返す関数
func RunLengthEncode(S string) []RL {
	N := len(S)
	res := make([]RL, 0)
	var l int
	for l < N {
		r := l
		for r < N && S[l] == S[r] {
			r++
		}
		res = append(res, RL{S[l], r - l})
		l = r
	}
	return res
}

// RunLengthDecode はRunlengthEncodeを復元する関数
func RunLengthDecode(EncodedRL []RL) string {
	res := make([]byte, 0)
	for _, v := range EncodedRL {
		for i := 0; i < v.Num; i++ {
			res = append(res, v.Str)
		}
	}
	return string(res)
}