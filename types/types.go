package types

import "time"

//easyjson:json
type SensorReadings struct {
	Created  time.Time `json:"created"`
	Bme280	*Bme280Readings
	Rainfall *TippingBucket
}

type TippingBucket struct {
	Inches float64 `json:"inches"`
	Milimeteres float32 `json:"milimeters"`
}

type Bme280Readings struct {
	TempF    float32 `json:"tempF"`
	Humidity float32 `json:"humidity"`
	Pressure float64 `json:"pressure"`
}

const (
	YYYYMMDD = "2022-06-01"
	TxInterval = time.Second * 10
)
