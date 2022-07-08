package convert

// tempCtoF Converts and returns Celsius in milidegrees to Fahrenheit degrees
func Celsius2Fahrenheit(tempC int32) (float32) {

	// degree := float32(tempC / 1000)
	// tempF := (degree * 1.8) + 32.0

	return (float32(tempC) * 9/5) + 32
}
