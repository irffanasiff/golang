package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var myString string
	// fmt.Scanln(&myString)
	// fmt.Println(myString)

	var name string = "irfan"
	var a, _ = fmt.Println(name)
	fmt.Println(a)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter your full name: ")
	myname, _ := reader.ReadString('\n')
	fmt.Println(myname)

}
