package main

import "fmt"

//* safe := true  ==> this does not work here because to declare in package scope things need keyword and short declarations does not have a keyword

func main() {

	safe, speed := true, 50
	name, secondName := "nikola", "tesla"
	nFiles, valid, msg := 10, true, "hello"

	fmt.Println(
		safe, speed, name, secondName, nFiles, valid, msg,
	)

	//%   redeclaration
	//? at least one variable in redeclaration should be a new variable to unless redeclaration wont work

	var isAlive bool = false
	isAlive, fullName := true, "Jon Procer"

	fmt.Println(isAlive, fullName)

	//! => if you dont know the intial value use variable declaration else use short declaration

	//? => if you need a package scoped variable use variable declaration

	//* => if you want to group variables together for greater readability use variable declaration

	var (
		video string

		duration, current int
	)

	//#  => if you want to redeclare and keep the code consize use short declarations

	width, height := 100, 50

	width, color := 50, "red"

	fmt.Println(
		video, duration, current, width, height, color,
	)

}
