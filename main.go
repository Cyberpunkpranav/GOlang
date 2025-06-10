package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
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

// port := os.Getenv("PORT") // Get the port from Render environment
//
//	if port == "" {
//		port = "8000" // Use 8000 for local testing
//	}
//
// // todo.Todo()
// http.HandleFunc("/", handler)
// http.HandleFunc("/fizz-buzz", handlerfb)
// http.HandleFunc("/whatsapp/webhook", webhook.WhatsAppWebhook)
// err := http.ListenAndServe(":"+port, nil)
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
// func function1(value int) {
// 	log.Printf("Printing value of funtion2 in function1 : %d", value)
// }
// func function2(value int) {
// 	log.Printf("Printing value of funtion1 in function2 :  %d", value)
// }
// func main() {

// 	channel := make(chan int)
// 	go func() {
// 		channel <- 1
// 		channel <- 2
// 		close(channel)
// 	}()
// 	val1 := <-channel
// 	val2 := <-channel
// 	go function1(val2)
// 	go function2(val1)
// 	select {}

// }

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from goroutine!")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()
	// time.Sleep(time.Second) // wait for goroutine to finish
	fmt.Println("Main ends")
}
