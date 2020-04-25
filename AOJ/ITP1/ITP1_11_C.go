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

func main() {
	var Dice1, Dice2 dice
	var cnt int
	fmt.Scan(&Dice1.o, &Dice1.tw, &Dice1.th, &Dice1.fo, &Dice1.fi, &Dice1.s)
	fmt.Scan(&Dice2.o, &Dice2.tw, &Dice2.th, &Dice2.fo, &Dice2.fi, &Dice2.s)
	rotdice := map[int]dice{
		1: Dice2,
		2: Dice2.rot("N"),
		3: Dice2.rot("W"),
		4: Dice2.rot("E"),
		5: Dice2.rot("S"),
		6: Dice2.rot("N").rot("N"),
	}
	for i := 1; i <= 6; i++ {
		Dice2 = rotdice[i]
		for {
			if Dice1 == Dice2 {
				fmt.Println("Yes")
				return
			}
			if cnt == 3 {
				cnt = 0
				break
			}
			Dice2 = Dice2.rot("R")
			cnt++
		}
	}
	fmt.Println("No")
}
