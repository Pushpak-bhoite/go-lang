package main

import "fmt"

func main() {
	loginCount := 23
	var result string

	if loginCount > 25 {
		result = " more than 25 "
	} else {
		result = "less than 25"
	}

	if num := 3; num> 4{
		fmt.Println("nu is gret that 4")
	}else{
		fmt.Println("nu is less that 4")
	}
	fmt.Println("result->", result)
}
