package convert

// CalculateRain returns rainfall amount by multiplying the number of tips and size.
func CalculateRain(tips uint, bucketSize float64) float64 {

	// fmt.Printf("Rainfall in inches: %v\n", amount)
	return float64(tips) * bucketSize
}
