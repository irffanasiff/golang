package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var name string
	var userRating string

	//front end
	//take name as input
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your full name")
	name, _ =reader.ReadString('\n') //readString reads until the first occurenece of delim in the input, returning a string containing the data up to and including the delimiter
    
	//take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our resturant between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	mynumRating,_ := strconv.ParseFloat(strings.TrimSpace(userRating), 64) 
	
	//back end
	fmt.Printf("%v %v\n", name, mynumRating)
}