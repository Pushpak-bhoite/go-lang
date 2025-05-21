package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now())

	diceNumber := rand.Intn(2) + 1
	fmt.Println("diceNumber->", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("the no is one")
		fallthrough
	case 2:
		fmt.Println("the no is two")
	default:
		fmt.Println("the no is dewfaukt")
	}

}
