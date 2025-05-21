package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"tomato", "apple"}
	fmt.Printf("Typs is %T", fruitList)

	fruitList = append(fruitList, "banana", "milk", "mango")
	fmt.Println("appended :", fruitList)

	fmt.Println("slice 1st to end :", fruitList[1:])
	fmt.Println("slice 1st to 3 :", fruitList[1:3])
	fmt.Println("slice 0 to 3 :", fruitList[:3]) // default starts from 0

	// ************ make *************
	highScore := make([]int, 4)
	highScore[0] = 3
	highScore[1] = 2
	highScore[2] = 5
	highScore[3] = 4
	fmt.Println("highScore->", highScore)

	highScore = append(highScore, 7, 8, 9) // memory will be reallocated and expanded for this highScore varible, even its default size is 4( evetually all values will e accomodated =)
	fmt.Println("extended highScore : ", highScore)
	sort.Ints(highScore)
	fmt.Println("sort", highScore) // sorting and all methods are available in the slice not in array

	// *********** Remove index value ***********
	var courses = []string{"reactJs", "js", "switft", "ruby", "AI"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses->", courses)
}
