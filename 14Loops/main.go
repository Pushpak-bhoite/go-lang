package main

import "fmt"

func main() {
	weekDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 1; i < len(weekDays); i++ {
		fmt.Println(weekDays[i])
	}
	for i, value := range weekDays {
		fmt.Println(i, value)
	}

	rogueVal := 1
	for rogueVal < 7 {
		if rogueVal == 2 {
			goto hello
		}

		if rogueVal == 5 {
			rogueVal++
			continue
		}
		rogueVal++
		fmt.Println(rogueVal)
	}

hello: // this is kind of label
	fmt.Println("go to directed to label lco")
}
