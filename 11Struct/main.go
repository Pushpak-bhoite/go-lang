package main

import "fmt"

func main() {
	hitesh := User{"Hitesh", "demo@gmail.com", true, 16}
	fmt.Println("hitesh->", hitesh)
	fmt.Printf("hitesh-> %+v \n", hitesh) // * remember %+v gives you whole structure
	fmt.Printf("hitesh-> %+v", hitesh.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	age    int
}
