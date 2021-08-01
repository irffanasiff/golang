package main

import (
	"fmt"
)

func main() {
	a := 10             //1010 --> 10
	b := 3              //0011 --> 3
	fmt.Println(a & b)  //0010 --> 2
	fmt.Println(a | b)  //1011 --> 11
	fmt.Println(a ^ b)  //1001 --> 9
	fmt.Println(a &^ b) //0100 --> 8

	//bit shifting
	x := 8              //2^3
	fmt.Println(x << 3) //2^3 * 2^3 =2^6
	fmt.Println(x >> 3) //2^3 / 2^3 =2^0

}
