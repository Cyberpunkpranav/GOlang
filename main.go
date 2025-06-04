package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is listening")
}
func main() {
	fmt.Println("Hello world")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
