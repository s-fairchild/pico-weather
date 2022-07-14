package convert

func MilliPA2MilliBar(mPA int32) float64 {

	return float64(mPA) / 100000
}

// Mm2Inches converts millimeters to inches
func Mm2Inches(mm float32) float32 {

	return mm * 2.54
}

// tempCtoF Converts and returns Celsius in milidegrees to Fahrenheit degrees
func Celsius2Fahrenheit(tempC int32) (float32) {

	return (float32(tempC) * 9/5) + 32
}