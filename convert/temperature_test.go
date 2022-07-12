package convert

import (
	"fmt"
	"testing"
)

var (
	likelyTemps = []float32{
		// 0,
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
)

func TestCelsius2Fahrenheit(t *testing.T) {

	err := testTempRanges(likelyTemps, 1, t.Name())
	if err != nil {
		t.Fatal(err)
	}
}

func testTempRanges(controls []float32, interval int, name string) error {

	// var startInterval int
	for i:=0; i<len(likelyTemps); i++ {
		control := likelyTemps[i]
		tempF := Celsius2Fahrenheit(int32(i))
		// fmt.Printf("Comparing %v to %v\n", int32(i), control)
		if tempF != control {
			return fmt.Errorf("%v got %v, wants %v", name, tempF, control)
		}
	}
	return nil
}
