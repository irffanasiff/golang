package main

import "fmt"

func main(){
	a:= 6
	//fmt.Printf(a)
	fmt.Printf("%v = %T \n",a,a)
}

// <Printf> - "Print Formatter" this function allows you to format numbers, variables and strings into the first string parameter you give it. <Printf> formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.

//<Print> formats using the default formats for its operands and writes to standard output. Spaces are added between operands when neither is a string. <Print> - "Print" This cannot format anything, it simply takes a string and print it 

//<Println> formats using the default formats for its operands and writes to standard output. <Println> - "Print Line" same thing as Printf() however it will append a newline character\n