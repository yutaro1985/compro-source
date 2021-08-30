// FastPrimeFactorize osa_k法を使った高速素因数分解
// sieveを作るところのjはi*2スタートでもよい？
// https://qiita.com/rsk0315_h4x/items/ff3b542a4468679fb409
func FastPrimeFactorize(n int, sieve []int) map[int]int {
	pf := make(map[int]int)
	for n > 1 {
		pf[sieve[n]]++
		n /= sieve[n]
	}
	return pf
}

// MakeSieve は高速素数判定に使う篩を作る。
// Sieve[i] < i なら合成数なので、Sieveは素数判定にも使える
// ただし、Sieve[1] = 1 なので1の場合に注意
func MakeSieve(max int) []int {
	sieve := make([]int, max)
	for i := range sieve {
		sieve[i] = i
	}
	for i := 2; i*i <= max; i++ {
		if sieve[i] == i {
			for j := i; j < max; j += i {
				if sieve[j] == j {
					sieve[j] = i
				}
			}
		}
	}
	return sieve
}
