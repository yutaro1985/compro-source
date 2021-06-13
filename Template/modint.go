type mint int64

func (mi mint) mod() mint {
	mi %= mint(mod)
	if mi < 0 {
		mi += mint(mod)
	}
	return mi
}
func (mi mint) inv() mint {
	// https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#3-5-%E6%8B%A1%E5%BC%B5-euclid-%E3%81%AE%E4%BA%92%E9%99%A4%E6%B3%95%E3%81%AB%E3%82%88%E3%82%8B%E9%80%86%E5%85%83%E8%A8%88%E7%AE%97
	b, u, v := mint(mod), mint(1), mint(0)
	for b != 0 {
		t := mi / b
		mi -= t * b
		mi, b = b, mi
		u -= t * v
		u, v = v, u
	}
	u %= mint(mod)
	if u < 0 {
		u += mint(mod)
	}
	return u
	// return mi.pow(mint(0).sub(2))
}
func (mi mint) pow(n mint) mint {
	p := mint(1)
	for n > 0 {
		if n&1 == 1 {
			p.mulSelf(mi)
		}
		mi.mulSelf(mi)
		n >>= 1
	}
	return p
}
func (mi mint) factorial() mint {
	// var n mint
	if mi < 0 {
		fmt.Println("[ERROR] Input nums 0 or greater.")
		os.Exit(1)
	}
	if mi == 1 || mi == 0 {
		return 1
	}
	return (mi - 1).factorial().mul(mi)
}
func (mi mint) add(n mint) mint {
	return (mi + n).mod()
}
func (mi mint) sub(n mint) mint {
	return (mi - n).mod()
}
func (mi mint) mul(n mint) mint {
	return (mi * n).mod()
}
func (mi mint) div(n mint) mint {
	return mi.mul(n.inv())
}
func (mi *mint) addSelf(n mint) *mint {
	*mi = mi.add(n)
	return mi
}
func (mi *mint) subSelf(n mint) *mint {
	*mi = mi.sub(n)
	return mi
}
func (mi *mint) mulSelf(n mint) *mint {
	*mi = mi.mul(n)
	return mi
}
func (mi *mint) divSelf(n mint) *mint {
	*mi = mi.div(n)
	return mi
}

// SingleCOM はmodintを使ってO(K)で一つの二項係数を求める関数
func SingleCOM(N, K int) mint {
	if N < 0 || K < 0 || K > N {
		return 0
	}
	res := mint(1)
	for i := 0; i < K; i++ {
		res.mulSelf(mint(N - i))
		res.divSelf(mint(i + 1))
	}
	return res
}

// modCOMで使うスライス
var fac = make([]mint, maxsize)
var finv = make([]mint, maxsize)
var inv = make([]mint, maxsize)

// COMinit で COMで使うためのテーブルを作る前処理
// O(N)
// https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#5-%E4%BA%8C%E9%A0%85%E4%BF%82%E6%95%B0-ncr
func ModCOMinit() {
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < maxsize; i++ {
		fac[i] = fac[i-1].mul(mint(i))
		inv[i] = mint(mod) - inv[mod%i].mul(mint(mod/i))
		finv[i] = finv[i-1].mul(inv[i])
	}
}

// FascModCOM nCkを求める。COMinitを先に実行する
// COMinitの結果を使ってO(1)で行える
func FastModCOM(n, k int) mint {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n].mul(finv[k].mul(finv[n-k]))
}
