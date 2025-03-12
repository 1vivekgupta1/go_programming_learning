package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Super Secret Information")
}
func isAuthorize(endpoint func(http.ResponseWriter,*http.Request))


func handleRequests(){
	http.Handle("/",isAuthorize(homepage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main()  {
	fmt.Printf("Server")
	handleRequests()
}
