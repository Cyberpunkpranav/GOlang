package webhooks

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func WhatsAppWebhook(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	log.Println("📥 Incoming webhook payload:")
	log.Println(string(bodyBytes))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)

	if mode == "subscribe" && token == "practicewebhoook" {
		fmt.Fprint(w, challenge)
		log.Println("Webhook verified successfully.")
	}
}
