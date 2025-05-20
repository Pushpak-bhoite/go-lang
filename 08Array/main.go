package main

import "fmt"

func main() {

	var fruits [4]string

	fruits[0] = "tomato"
	fruits[1] = "mango"
	fruits[3] = "banana"

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var veggies = [4]string{"chilli", "milli", "pilli"};
	fmt.Println(veggies)

}
