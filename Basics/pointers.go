package main

import (
	"fmt"
)

func main() {
	// & adress of that variable | * returns the value at that adress it is a derefrencing operator
	// var x int = 1
	// var y int
	// var ip *int //ip is pointer to int
	// ip = &x     //ip now points to x
	// y = *ip     //y is now 1
	// fmt.Println(x, y, ip, *ip, &x, &y, &ip)

	// new alternate way to create a variable
	// new() function creates a variable and returns a pointer to the varialbe
	// variable is initialized to zero

	ptr := new(int) //same as var ptr * int and then *ptr = 3
	*ptr = 3
	fmt.Println(ptr, *ptr, &ptr)
}
