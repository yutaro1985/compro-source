package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	Weathers := []string{"Sunny", "Cloudy", "Rainy"}
	for i, weather := range Weathers {
		if weather == "Rainy" {
			fmt.Println("Sunny")
			return
		} else {
			if weather == S {
				fmt.Println(Weathers[i+1])
				return
			}
		}
	}
}
