package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/risk", risk)
	http.ListenAndServe(":8090", nil)
}
