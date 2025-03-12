package main

import (
	"fmt"
	"io"

	"os"
)

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}

func readFile(filname string){
	databyte, err := os.ReadFile(filname)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}


func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil{
	// 	panic(err)
	// }

	checkNilErr(err)

	length , err := io.WriteString(file,content)
	checkNilErr(err)
	fmt.Println("Length is : ",length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}