package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Fruit list is: ",len(fruitList))

	var vegList = [5]string{"potato","beans","mushroom"}

	fmt.Println("Vegy list is:", len(vegList))
}
