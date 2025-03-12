package main

import (
	"fmt"
	"os"
)
var MySigningKey= []byte(os.Getenv("SECRET_KEY"))
func handleRequests(w http.ResponseWriter, r *http.Request) {
	
}
func main() {
	handleRequests()
}