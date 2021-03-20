package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Order struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type OrderResponse struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type OrderPayload struct {
	Data Order
}

type OrderPayloadResponse struct {
	Data OrderResponse
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received home request\n")
	w.Write([]byte(fmt.Sprintf("Hello from Golang!!!")))
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for ping\n")
	w.Write([]byte("pong"))
}

func handlerOrderSubmit(w http.ResponseWriter, r *http.Request) {

	var payload OrderPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	log.Printf("payload data.Id: %s\n", payload.Data.Id)
	log.Printf("payload data.Description: %s\n", payload.Data.Description)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var payloadResponse OrderPayloadResponse
	payloadResponse.Data.Id = payload.Data.Id
	payloadResponse.Data.Description = payload.Data.Description
	payloadResponse.Data.Status = "Accepted"

	defer r.Body.Close()

	respondWithJSON(w, http.StatusCreated, payloadResponse)
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handlerHome)
	r.HandleFunc("/ping", handlerPing)
	r.HandleFunc("/order/submit", handlerOrderSubmit).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":5500",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Configure Logging
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
