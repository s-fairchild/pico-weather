package types

import "time"

//easyjson:json
type SensorReadings struct {
	Created  time.Time `json:"created"`
	Bme280	*Bme280Readings
	Rain *TippingBucket
}

type TippingBucket struct {
	Amount float64 `json:"inches"`
	// Millimeters float64 `json:"milometers"`
}

type Bme280Readings struct {
	Temp    float32 `json:"tempF"`
	Humidity float32 `json:"humidity"`
	Pressure float64 `json:"pressure"`
}

const (
	YYYYMMDD = "2022-06-01"
	TxInterval = time.Second * 5
)
