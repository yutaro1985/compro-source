package main

import "fmt"

type dice struct {
	o, tw, th, fo, fi, s int
}

func (Dice dice) rot(D string) dice {
	for _, d := range D {
		switch d {
		case 'N':
			Dice.o, Dice.tw, Dice.fi, Dice.s = Dice.tw, Dice.s, Dice.o, Dice.fi
		case 'E':
			Dice.o, Dice.th, Dice.fo, Dice.s = Dice.fo, Dice.o, Dice.s, Dice.th
		case 'W':
			Dice.o, Dice.th, Dice.fo, Dice.s = Dice.th, Dice.s, Dice.o, Dice.fo
		case 'S':
			Dice.o, Dice.tw, Dice.fi, Dice.s = Dice.fi, Dice.o, Dice.s, Dice.tw
		case 'R':
			Dice.tw, Dice.th, Dice.fo, Dice.fi = Dice.th, Dice.fi, Dice.tw, Dice.fo
		case 'L':
			Dice.tw, Dice.th, Dice.fo, Dice.fi = Dice.fo, Dice.tw, Dice.fi, Dice.th
		}
	}
	return Dice
}

func (Dice dice) isEqual(D dice) bool {
	var cnt int
	rotdice := map[int]dice{
		1: D,
		2: D.rot("N"),
		3: D.rot("W"),
		4: D.rot("E"),
		5: D.rot("S"),
		6: D.rot("N").rot("N"),
	}
	for i := 1; i <= 6; i++ {
		D = rotdice[i]
		for {
			if Dice == D {
				return true
			}
			if cnt == 3 {
				cnt = 0
				break
			}
			D = D.rot("R")
			cnt++
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	var Dice1 dice
	fmt.Scan(&Dice1.o, &Dice1.tw, &Dice1.th, &Dice1.fo, &Dice1.fi, &Dice1.s)
	for i := 0; i < n-1; i++ {
		var Dice dice
		fmt.Scan(&Dice.o, &Dice.tw, &Dice.th, &Dice.fo, &Dice.fi, &Dice.s)
		if Dice1.isEqual(Dice) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
