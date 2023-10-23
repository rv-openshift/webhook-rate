package main

import (
	"encoding/json"
	"fmt"
	"webhook-rate/mail"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func writeConsole(writer http.ResponseWriter, content, status, body string, fail bool) error  {
	if !fail {
		writer.Header().Set("Content-Type", content)
		writer.WriteHeader(http.StatusOK)	
	} else {
		writer.WriteHeader(http.StatusTooManyRequests)		
	}
	message := Message{
		Status: status,
		Body:   body,
	}
	err := json.NewEncoder(writer).Encode(&message)	
	return err
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/ping" {
		log.Println(writer, "Invalid request method.", 405)
		return
	}
	if request.Method == "GET" {
		log.Println("Webhook is running...(GET)")
		return
	}
	err := writeConsole(writer, "application/json", "Successful", "Hi! You've reached the API. How may I help you?", false)	
	if err != nil {
		return
	} else {
		er := mail.SendmaiL("venerayan@gmail.com", os.Args[1]) // os-args is the app-password of gmail
		if er != nil {
			fmt.Println("Email have error: %v", er)
		}
		return
	}
}

func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 1)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			fmt.Println("Request Failed.  The API is at capacity, try again later.")
			writeConsole(w, "", "Request Failed", "The API is at capacity, try again later.", true)				
			return
		} else {
			next(w, r)
		}
	})
}

func main() {
	log.Println("Webhook is running...")
	http.Handle("/ping", rateLimiter(endpointHandler))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Println("There was an error listening on port :9000", err)
	}
}
