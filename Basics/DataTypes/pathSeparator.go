package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println("Direcotry: ", dir, "\nfile Name: ", file)

	var dir2 string
	dir2, _ = path.Split("css/main.css") //using blank identifier
	fmt.Println("Direcotry: ", dir2)
 
	//more consize code
	dir3, file3 := path.Split("css/main.css")
	fmt.Println("Direcotry: ", dir3, "\nfile Name: ", file3)


}
