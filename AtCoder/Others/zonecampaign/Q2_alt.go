package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, initialBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

// 問題によって値は調整する
const (
	mod     = int(1e9) + 7
	maxsize = 510000
	INF     = 1 << 60
)

func main() {
	var ans int
	isPrime := map[int]bool{
		2:   true,
		3:   true,
		5:   true,
		7:   true,
		11:  true,
		13:  true,
		17:  true,
		19:  true,
		23:  true,
		29:  true,
		31:  true,
		37:  true,
		41:  true,
		43:  true,
		47:  true,
		53:  true,
		59:  true,
		61:  true,
		67:  true,
		71:  true,
		73:  true,
		79:  true,
		83:  true,
		89:  true,
		97:  true,
		101: true,
		103: true,
		107: true,
		109: true,
		113: true,
		127: true,
		131: true,
		137: true,
		139: true,
		149: true,
		151: true,
		157: true,
		163: true,
		167: true,
		173: true,
		179: true,
		181: true,
		191: true,
		193: true,
		197: true,
		199: true,
		211: true,
		223: true,
		227: true,
		229: true,
		233: true,
		239: true,
		241: true,
		251: true,
		257: true,
		263: true,
		269: true,
		271: true,
		277: true,
		281: true,
		283: true,
		293: true,
		307: true,
		311: true,
		313: true,
		317: true,
		331: true,
		337: true,
		347: true,
		349: true,
		353: true,
		359: true,
		367: true,
		373: true,
		379: true,
		383: true,
		389: true,
		397: true,
		401: true,
		409: true,
		419: true,
		421: true,
		431: true,
		433: true,
		439: true,
		443: true,
		449: true,
		457: true,
		461: true,
		463: true,
		467: true,
		479: true,
		487: true,
		491: true,
		499: true,
		503: true,
		509: true,
		521: true,
		523: true,
		541: true,
		547: true,
		557: true,
		563: true,
		569: true,
		571: true,
		577: true,
		587: true,
		593: true,
		599: true,
		601: true,
		607: true,
		613: true,
		617: true,
		619: true,
		631: true,
		641: true,
		643: true,
		647: true,
		653: true,
		659: true,
		661: true,
		673: true,
		677: true,
		683: true,
		691: true,
		701: true,
		709: true,
		719: true,
		727: true,
		733: true,
		739: true,
		743: true,
		751: true,
		757: true,
		761: true,
		769: true,
		773: true,
		787: true,
		797: true,
		809: true,
		811: true,
		821: true,
		823: true,
		827: true,
		829: true,
		839: true,
		853: true,
		857: true,
		859: true,
		863: true,
		877: true,
		881: true,
		883: true,
		887: true,
		907: true,
		911: true,
		919: true,
		929: true,
		937: true,
		941: true,
		947: true,
		953: true,
		967: true,
		971: true,
		977: true,
		983: true,
		991: true,
		997: true,
	}
	for i := 0; i < 20*20; i++ {
		if isPrime[nextInt()] {
			ans++
		}
	}
	fmt.Println(ans)
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
