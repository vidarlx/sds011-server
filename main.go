package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server started on 9999")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}

	fmt.Println("Server started on 9999")
}
