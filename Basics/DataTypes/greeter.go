package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("%#v\n", os.Args)

	fmt.Println("path: ", os.Args[0])
	fmt.Println("Args[1]: ", os.Args[1])
	fmt.Println("Args[2]: ", os.Args[2],
		len(os.Args))

}
