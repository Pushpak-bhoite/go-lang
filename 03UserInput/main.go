// bufio package provides input functionality - cheput on go doc
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("the input is : ", input)
	fmt.Printf("the input is : %T", input)
}
