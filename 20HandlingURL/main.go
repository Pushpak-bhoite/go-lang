package main

import (
	"fmt"
	"net/url"
)

func main() {

	myUrl := "https://api.agify.io:3000/?name=michael&country_id=US"

	fmt.Println("hello->", myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	fmt.Println(result.Query()) // gives value in map

	// ******* Create Url **********
	partsOfUrl := &url.URL{ //& - is must to use
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/test",
		RawPath: "user",
	}
	fmt.Println(partsOfUrl)
	fmt.Println(partsOfUrl.String())
}
