package main

import "fmt"

func main() {
	langs := make(map[string]string)
	langs["JS"] = "JavaScript"
	langs["RB"] = "Ruby"
	langs["PY"] = "Python"

	for key, value := range langs {
		fmt.Printf("For key v, value is %v ,key is :%v\n", key, value)
	}

	for key, value := range langs {
		fmt.Println("key: ", key, "value: ", value)
	}
	fmt.Println("map ->", langs)

}
