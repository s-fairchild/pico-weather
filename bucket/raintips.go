package bucket

var Tips uint

// CalculateRain returns rainfall amount by multiplying the number of tips and size.
func calculateRain(tips uint, bucketSize float64) float64 {

	return float64(tips) * bucketSize
}
