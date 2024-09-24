package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const targetURL = "https://api.ssms-pilani.in/menu-next-seven-days" // Replace with your desired URL

var (
	lastResponse string
	lastFetch    time.Time
	mutex        sync.Mutex
)

func fetchData() error {
	resp, err := http.Get(targetURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()
	lastResponse = string(body)
	lastFetch = time.Now()
	return nil
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	if err := fetchData(); err != nil {
		http.Error(w, "Failed to refresh data", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Data refreshed successfully")
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	if lastResponse == "" {
		http.Error(w, "No data available. Please refresh first.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, lastResponse)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is healthy")
}

func main() {
	http.HandleFunc("/refresh", refreshHandler)
	http.HandleFunc("/fetch", fetchHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
