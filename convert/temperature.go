package convert

// tempCtoF Converts and returns Celsius in milidegrees to Fahrenheit degrees
func Celsius2Fahrenheit(tempC int32) (float32) {

	return (float32(tempC) * 9/5) + 32
}
