// Crt は中国剰余定理の実装
// common_mathを一緒に使用すること。
// 答えがない場合は(0,0)が返る
// https://github.com/monkukui/ac-library-go/blob/master/math/math.go
func Crt(r, m []int64) (int64, int64) {
	if len(r) != len(m) {
		panic("")
	}
	n := len(r)
	r0 := int64(0)
	m0 := int64(1)
	for i := 0; i < n; i++ {
		if !(1 <= m[i]) {
			panic("")
		}
		r1 := SafeMod(r[i], m[i])
		m1 := m[i]
		if m0 < m1 {
			r0, r1 = r1, r0
			m0, m1 = m1, m0
		}
		if m0%m1 == 0 {
			if r0%m1 != r1 {
				return 0, 0
			}
			continue
		}
		g, im := InvGcd(m0, m1)

		u1 := m1 / g
		if (r1-r0)%g != 0 {
			return 0, 0
		}

		x := (r1 - r0) / g % u1 * im % u1

		r0 += x * m0
		m0 *= u1
		if r0 < 0 {
			r0 += m0
		}
	}
	return r0, m0
}