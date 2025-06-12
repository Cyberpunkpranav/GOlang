package main

import (
	"fmt"
	pgDB "go-server/database"
	webhook "go-server/webhooks"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server on 8080")
	log.Println("server is listening")
}

func main() {
	appEnv := "local"
	// appEnv := os.Getenv("APP_ENV")
	http.HandleFunc("/whatsapp/webhook", webhook.WhatsAppWebhook)

	pgDB.ConnectPostgreSql()
	http.HandleFunc("/", handler)
	if appEnv == "local" {
		http.ListenAndServe(":8080", nil)
	}
	if appEnv == "docker" {
		http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil)
	}
}
