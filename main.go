package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type sensor_data struct {
	date string
}

func dataHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t sensor_data
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	log.Println(t.date)
}

func main() {
	http.HandleFunc("/data", dataHandler)

	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
