package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato","Peach"}

	fruitList = append(fruitList, "Mango","Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	// slices with inplace
	fmt.Println(fruitList)


	highScores := make([]int,4)

	highScores[0] = 234
	highScores[1] = 606
	highScores[2] = 347
	highScores[3] = 857
	// highScores[4] = 578

	highScores = append(highScores, 555,666,321)
	fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	
}