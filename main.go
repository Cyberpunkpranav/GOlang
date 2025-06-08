package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	webhook "go-server/webhooks"
)

func FizzBuzz(n int) map[int]string {
	var i int
	result := make(map[int]string)
	for i = 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			// fmt.Println("FizzBuzz")
			result[i] = "FizzBuzz"
			continue
		} else if i%3 == 0 {
			// fmt.Println("Fizz")
			result[i] = "Fizz"
			continue
		} else if i%5 == 0 {
			// fmt.Println("Buzz")
			result[i] = "Buzz"
			continue
		}
		// fmt.Println(i)
		result[i] = strconv.Itoa(i)
	}
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is listening")
}

func handlerfb(w http.ResponseWriter, r *http.Request) {
	n := 100
	result := FizzBuzz(n)
	response := struct {
		Status string         `json:"status"`
		Data   map[int]string `json:"data"`
	}{
		Status: "200",
		Data:   result,
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(string(jsonData))
	w.Write(jsonData)
}

func main() {
	port := os.Getenv("PORT") // Get the port from Render environment
	if port == "" {
		port = "8000" // Use 8000 for local testing
	}
	// todo.Todo()
	http.HandleFunc("/", handler)
	http.HandleFunc("/fizz-buzz", handlerfb)
	http.HandleFunc("/whatsapp/webhook", webhook.WhatsAppWebhook)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}
