package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

// AOJではgodsが使えないので通せない

func main() {
	stack := lls.New()

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	formula := strings.Split(sc.Text(), " ")
	for _, v := range formula {
		switch v {
		case "+":
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			stack.Push(a.(int) + b.(int))
		case "-":
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			stack.Push(a.(int) - b.(int))
		case "*":
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			stack.Push(a.(int) * b.(int))
		default:
			num, _ := strconv.Atoi(v)
			stack.Push(num)
		}
	}
	ans, _ := stack.Pop()
	fmt.Println(ans.(int))
}
