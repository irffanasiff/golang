package main

import(
	"fmt"
)

func main(){
//	var n bool
//	i := 1 == 1
//	m := 1 == 2
	
//   fmt.Printf("%v, %T\n", n, n)
//	fmt.Printf("%v, %T\n", i, i)
//	fmt.Printf("%v, %T\n", m, m)
//    var n uint16 = 42
//  fmt.Printf("%v, %T\n", n, n)
    a:= 10 
	b:=3
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)
}
/*  Printf - "Print Formatter" this function allows you to format numbers, variables and strings into the first string parameter you give it

Print - "Print" This cannot format anything, it simply takes a string and print it

Println - "Print Line" same thing as Printf() however it will append a newline character \n 

Println also inserts spaces between arguments. Print only inserts spaces between arguments when either argument is not a string.*/