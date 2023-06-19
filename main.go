package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API is running")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	errors := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	}

	// select random error
	rand.Seed(time.Now().UnixNano())
	w.WriteHeader(errors[rand.Intn(len(errors))])
}

func loadMemoryHandler(w http.ResponseWriter, r *http.Request) {
	// Create a byte slice with 1 GB
	_ = make([]byte, 1<<30)

	fmt.Fprint(w, "Increased memory load")
}

func loadCPUHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan int)

	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}
		}
	}()

	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	fmt.Fprint(w, "Increased CPU load")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/loadMemory", loadMemoryHandler)
	http.HandleFunc("/loadCpu", loadCPUHandler)

	http.ListenAndServe(":8080", nil)
}
