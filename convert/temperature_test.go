package convert

import (
	"testing"
)

var (
	likelyTemps = []float32{
		32,
		33.8, // 1 °C
		35.6,
		37.4,
		39.2,
		41.0,
		42.8,
		44.6,
		46.4,
		48.2,
		50.0,
		51.8,
		53.6,
		55.4,
		57.2,
		59.0,
		60.8,
		62.6,
		64.4,
		66.2,
		68.0,
		69.8,
		71.6,
		73.4,
		75.2,
		77.0,
		78.8,
		80.6,
		82.4,
		84.2,
		86.0,
		87.8,
		89.6,
		91.4,
		93.2,
		95.0,
		96.8,
		98.6,
		100.4,
		102.2,
		104.0,
	}
	highTemps = []float32{
		122.0,
		140.0,
		158.0,
		176.0,
		194.0,
		212.0, // 100 °C
		392.0,
		572.0,
		752.0,
		932.0,
		1112.0,
		1292.0,
		1472.0,
		1652.0,
		1832.0,
	}
)

func TestCelsius2Fahrenheit(t *testing.T) {

	const (
		zero = 0
		parity = -40
	)

	tempF := Celsius2Fahrenheit(parity)
	if tempF != parity {
		t.Fatalf("%v got %v, wants %v", t.Name(), tempF, parity)
	}

	testTempRanges(likelyTemps, 1)
	testTempRanges(highTemps, 10)
}

func testTempRanges(controls []float32, interval int) {

	t := testing.T{}

	for i:=0; i<len(likelyTemps); i++ {
		control := likelyTemps[i]
		tempF := Celsius2Fahrenheit(int32(i))
		t.Logf("Comparing %v to %v", int32(i), control)
		if tempF != control {
			t.Fatalf("%v got %v, wants %v", t.Name(), tempF, control)
			return 
		}
	}
}
