package main

import (
	"fmt"
	"net/http"
)

// this is main handle all function and listen to 8080 port
func main() {
	fmt.Print("server is running:\n")

	http.HandleFunc("/home", home)
	http.HandleFunc("/pageOne", pageOne)
	http.HandleFunc("/aboutUs", aboutUs)
	_ = http.ListenAndServe(":8080", nil)

}
