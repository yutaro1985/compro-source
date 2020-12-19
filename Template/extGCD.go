// extGCD は拡張ユークリッドの互除法の実装
// var x,y int を定義して渡すとax+by=GCD なるa,bのGCDとx,yを返してくれる
func extGCD(a, b int, x, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	d := extGCD(b, a%b, y, x)
	*y -= (a / b) * *x
	return d
}
