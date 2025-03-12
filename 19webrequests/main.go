package main

import (
	"fmt"
	"io"
	
	"net/http"
)

const url = "https://google.co.in"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err!= nil{
		panic(err)
	}
	fmt.Printf("Response is of Type :%T\n",response)
	defer response.Body.Close() // Caller's responsiblity to cloase the connection

	databytes, err := io.ReadAll(response.Body)

	if err!= nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
