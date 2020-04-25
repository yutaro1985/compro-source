package main

import "fmt"

type dice struct {
	o, tw, th, fo, fi, s int
}

func main() {
	var Dice dice
	fmt.Scan(&Dice.o, &Dice.tw, &Dice.th, &Dice.fo, &Dice.fi, &Dice.s)
	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var t, f int
		fmt.Scan(&t, &f)
		switch t {
		case Dice.tw:
			Dice = rot(Dice, "N")
		case Dice.th:
			Dice = rot(Dice, "W")
		case Dice.fo:
			Dice = rot(Dice, "E")
		case Dice.fi:
			Dice = rot(Dice, "S")
		case Dice.s:
			Dice = rot(rot(Dice, "N"), "N")
		}
		for Dice.tw != f {
			Dice = rot(Dice, "R")
		}
		fmt.Println(Dice.th)
	}
}

func rot(Dice dice, D string) dice {
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
