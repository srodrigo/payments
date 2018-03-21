package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8000
	log.Println(fmt.Sprintf("Listening on port %d", port))

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
