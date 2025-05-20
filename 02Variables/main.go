package main

import "fmt"

func main() {
	var username string = "Hitesh"
	fmt.Println(username)
	fmt.Printf("username is of type : %T \n", username)

	var age int
	fmt.Print("lets see default value of age : ", age)
	myNum := 34.90
	fmt.Print("\n var without declaring myNum : ", myNum)
}
