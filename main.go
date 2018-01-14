package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Properties struct {
	ValueType string `json:"value_type"`
	Value     string `json:"value"`
}

type SensorReads struct {
	SensorID   string       `json:"esp8266id"`
	SensorData []Properties `json:"sensordatavalues"`
}

type Record struct {
	PM25        float64
	PM10        float64
	Temperature float64
	Humidity    float64
	Date        time.Time
}

func dataHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	var t SensorReads
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	var r Record
	r.Date = time.Now()
	for _, v := range t.SensorData {
		switch v.ValueType {
		case "SDS_P1":
			val, e := strconv.ParseFloat(v.Value, 64)
			if e != nil {
				fmt.Println(e)
			}

			r.PM10 = val
			break

		case "SDS_P2":
			val, e := strconv.ParseFloat(v.Value, 32)
			if e != nil {
				fmt.Println(e)
			}
			r.PM25 = val
			break

		case "temperature":
			val, e := strconv.ParseFloat(v.Value, 32)
			if e != nil {
				fmt.Println(e)
			}
			r.Temperature = val
			break
		case "humidity":
			val, e := strconv.ParseFloat(v.Value, 32)
			if e != nil {
				fmt.Println(e)
			}
			r.Humidity = val
			break
		}
		//fmt.Println(k, v.ValueType, v.Value)
	}
	fmt.Println(r)
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
