package main

import (
	"net/http"
	"simple-webhook/src/api"
)

func main() {
	http.HandleFunc("/api/v1/servicenow", api.SnowHandler)

	http.ListenAndServe(":5001", nil)
}
