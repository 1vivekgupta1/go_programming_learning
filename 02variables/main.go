package main

import "fmt"

const LoginToken string = "gghghgh" //Varialbes declared in capital letters are public can be accessed by anyone file in a project.

func main() {
	var username string = "vivek"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n",username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variables is of type: %T \n",username)

	var smallVal uint8 = 255 //unsigned integers.
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n", smallVal)


	var smallFloat float64 = 255.45554511254451886
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type: %T \n", smallFloat)


	//Default values and some aliases
	var anotherVariables int //intial value is zero
	fmt.Println(anotherVariables)
	fmt.Printf("Variables is of type: %T \n", anotherVariables)

	//implicit type(means not exact or inferred type)

	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style

	numberofUser := 30000.0// := walrus operator
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variables is of type: %T \n", LoginToken)

}