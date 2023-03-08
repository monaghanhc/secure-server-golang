package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    // Define your API endpoint
    http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
        // Handle your API request logic here
        responseData := map[string]string{"message": "Hello, World!"}
        json.NewEncoder(w).Encode(responseData)
    })

    // Create a secure HTTPS server
    err := http.ListenAndServeTLS(":3000", "path/to/cert.pem", "path/to/key.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServeTLS: ", err)
    }
}
