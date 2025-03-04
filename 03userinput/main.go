package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"


	fmt.Println(welcome)


    //Initializing the reader class
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err err for try and catch functionalities

	input, _ := reader.ReadString('\n') // insider brackets delimiter is being specified

	fmt.Println("Thanks for rating , ", input)
	fmt.Printf("Type of this rating is %T",input)
}