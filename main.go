package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldWithTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	response := fmt.Sprintf("Hello, World! Current Time: %s", currentTime)
	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/", helloWorldWithTimeHandler)
	http.ListenAndServe(":8080", nil)
}
