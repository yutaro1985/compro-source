package main

import "fmt"

type dice struct {
	o, tw, th, fo, fi, s int
}

func main() {
	var Dice dice
	fmt.Scan(&Dice.o, &Dice.tw, &Dice.th, &Dice.fo, &Dice.fi, &Dice.s)
	var D string
	fmt.Scan(&D)
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
		}
	}
	fmt.Println(Dice.o)
}
