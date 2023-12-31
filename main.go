package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "API is running",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	// Set log flags to 0 to disable date and time prefix
	log.SetFlags(0)

	// 任意のステータスエラーをランダム生成
	rand.Seed(time.Now().UnixNano())
	errorIndex := rand.Intn(len(errors))
	w.WriteHeader(errors[errorIndex])

	response := Response{
		Message: http.StatusText(errors[errorIndex]),
	}

	// ログ出力をjson形式に変換する
	logData := map[string]interface{}{
		"timestamp":   time.Now().Format(time.RFC3339),
		"status_code": errors[errorIndex],
		"message":     response.Message,
	}
	logJSON, err := json.Marshal(logData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(string(logJSON))

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadMemoryHandler(w http.ResponseWriter, r *http.Request) {
	// 1ビットを29ビット左にシフト（2の29乗）で約500MBバイトのメモリ容量を表し、それを5倍することで約2.5GBのメモリ容量を確保
	mem := make([]byte, 5<<29)

	// 30秒後にメモリを解放する
	go func(memPtr *[]byte) {
		time.Sleep(30 * time.Second)
		*memPtr = nil // メモリへの参照を切り、ガーベジコレクションが解放できるようにする
	}(&mem)

	response := Response{
		Message: "Increased memory load",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadCPUHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan int)

	go func() {
		// 無限ループ作成
		for {
			select {
			// doneチャネルからメッセージが送られてきたとき（またはチャネルが閉じられたとき）無限ループを抜ける
			case <-done:
				return
			// 何もせず次のループへ進む
			default:
			}
		}
	}()

	// 10秒後にチャネルを閉じる
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	response := Response{
		Message: "Increased CPU load",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func delayedResponseHandler(w http.ResponseWriter, r *http.Request) {
	// 3秒遅延
	time.Sleep(3 * time.Second)

	response := Response{
		Message: "Delay occurred",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/loadMemory", loadMemoryHandler)
	http.HandleFunc("/loadCpu", loadCPUHandler)
	http.HandleFunc("/delayedResponse", delayedResponseHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
