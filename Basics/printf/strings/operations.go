package main

import "fmt"

func main() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2
	fmt.Println(average)

	//when using float result can be inaccurate

	var ratio float64 = 3 / 2
	fmt.Println(ratio)
	ratio = float64(int(3) / int(2))
	fmt.Println(ratio)
	ratio = float64(3) / 2
	fmt.Println(ratio)
	ratio = 3.0 / 2
	fmt.Println(ratio)
}
