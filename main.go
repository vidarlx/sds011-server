package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Properties struct {
	ValueType string `json:"value_type"`
	Value     string `json:"value"`
}

type SensorReads struct {
	SensorID   string       `json:"esp8266id"`
	SensorData []Properties `json:"sensordatavalues"`
}

func dataHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var t SensorReads
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	for k, v := range t.SensorData {
		fmt.Println(k, v.ValueType, v.Value)
	}
	// fmt.Println(t)
}

func main() {
	http.HandleFunc("/data", dataHandler)
	fmt.Println("Server started on 9099")
	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}

	fmt.Println("Server started on 9099")
}
