package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("response->%T\n", response)
	fmt.Println("response->", response)

	defer response.Body.Close() // callers responsiblity

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("hello->",content)
}
