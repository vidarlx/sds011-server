package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env"

	"github.com/ghanto/sds011-server/config"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Server started on 9999")

	if err := env.Parse(&config.ApplicationConfig); err != nil {
		fmt.Printf("%+v\n", err)
	}

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}

	fmt.Println("Server started on 9999")
}
