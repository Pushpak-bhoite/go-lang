// panic - its just a way to show errors and stop exection
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "hello im going in the file"

	file, err := os.Create("./tempFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is ", length)
	defer file.Close() // its a standard to use defer for close method as we might want to write code further and then execute close

	readFile("./tempFile.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println("databytes :->",databyte)

}
