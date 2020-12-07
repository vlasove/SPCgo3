//main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web!")
}

func hiWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from my go application!")
}

func main() {
	http.HandleFunc("/", helloWeb)
	http.HandleFunc("/hi", hiWeb)
	fmt.Println("Our application working......")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
