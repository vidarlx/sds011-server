package sensor

import "time"

type properties struct {
	ValueType string `json:"value_type"`
	Value     string `json:"value"`
}

type sensorReads struct {
	SensorID   string       `json:"esp8266id"`
	SensorData []properties `json:"sensordatavalues"`
}

// Record single measurement record
type Record struct {
	PM25        float64
	PM10        float64
	Temperature float64
	Humidity    float64
	Date        time.Time
}
