// Note: we cannt declare functions inside main func, we can only call them
package main

import "fmt"

func main() {
	greet()
	fmt.Println("")
	fmt.Println(addVal(3, 2))
	res, text := unawareOfParam(3, 4, 4)
	fmt.Println("haha->", res, text)
	fmt.Println(unawareOfParam(3, 4, 4))
}

func greet() {
	fmt.Println("Hello")
}

func addVal(valOne int, valTwo int) int { // rerturn signature mandatory
	return valOne + valTwo
}

func unawareOfParam(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total = total + val
	}
	return total, "im extram one text "
}
