package main

import "fmt"

func main() {
	hitesh := User{"Hitesh", "demo@gmail.com", true, 16}
	
	hitesh.GetEmail()
	hitesh.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){   // now this become method as we defined User struct in the  
	fmt.Println("user name is ", u.Email);
}

func (u User) GetEmail(){   // now this become method as we defined User struct in the  
	u.Email ="hello@gamil.com"
	fmt.Println("user name is ", u.Email);
}