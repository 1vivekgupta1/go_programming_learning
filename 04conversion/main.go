package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to our pizza app")

	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating , ", input)

	numRating, err := strconv.ParseFloat(stings.TrimSpace(input),64)
	
	if err!= nil{
		fmt.Println(err)
	}
	else {
		fmt.println("Added q to your raing:",numRatingnumRating)
	}
}