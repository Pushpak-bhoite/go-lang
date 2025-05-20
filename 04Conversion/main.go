package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Print("this is input ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// if you want to see error  - try by removing TrimSpace function from above 

	if err != nil {
		fmt.Println("Watch the error -> ", err)
	} else {
		fmt.Println("Added 1 to your rating -> ", numRating+1)
	}

}
