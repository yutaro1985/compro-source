// COMで使うスライス
var fac = make([]int, maxsize)
var finv = make([]int, maxsize)
var inv = make([]int, maxsize)

// COMinit で COMで使うためのテーブルを作る前処理
// O(N)
// https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#5-%E4%BA%8C%E9%A0%85%E4%BF%82%E6%95%B0-ncr
func COMinit() {
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < maxsize; i++ {
		fac[i] = fac[i-1] * i % mod
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
		finv[i] = finv[i-1] * inv[i] % mod
	}
}

// COM nCkを求める。COMinitを先に実行する
// COMinitの結果を使ってO(1)で行える
func COM(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % mod) % mod
}
