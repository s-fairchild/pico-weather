package convert

import (
	"testing"
	"fmt"
)

func TestMilliPa2MilliBar(t *testing.T) {

	var (
		controls = []float64{
			0,
			0.000010,
			0.000020,
			0.000030,
			0.000040,
			0.000050,
			0.000060,
			0.000070,
			0.000080,
			0.000090,
			0.00010,
			0.00011,
			0.00012,
			0.00013,
			0.00014,
			0.00015,
			0.00016,
			0.00017,
			0.00018,
		}

		seaLevel = 1013.25

		edgeCases = map[float64]float64{
			8.7e+7: 870, // lowest recorded pressure recorded in a typhoon in the Pacific Ocean.
			1.084e+8: 1084, // The highest air pressure recorded was 1084 mb in Siberia
		}
	)

	for i, v := range controls {
		n := MilliPA2MilliBar(int32(i))
		if n != v {
			t.Fatalf("got %v, expected %v", n, v)
		}
	}

	for i, v := range edgeCases {
		n := MilliPA2MilliBar(int32(i))
		if n != v {
			t.Fatalf("got %v, expected %v", n, v)
		}
	}

	n := MilliPA2MilliBar(101325000)
	if n != seaLevel {
		t.Fatalf("got %v, expected %v", n, seaLevel)
	}

}

func TestCalculateRain(t *testing.T) {

	bucketSizemm := 0.2794
	expect := 2.2352
	n := CalculateRain(8, bucketSizemm)
	if n != expect {
		t.Fatalf("%v expected %v got %v", t.Name(), expect, n)
	}
}


func TestCelsius2Fahrenheit(t *testing.T) {

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

	err := testTempRanges(likelyTemps, 1, t.Name())
	if err != nil {
		t.Fatal(err)
	}
}

func testTempRanges(controlTemps []float32, interval int, name string) error {

	// var startInterval int
	for i:=0; i<len(controlTemps); i++ {
		control := controlTemps[i]
		tempF := Celsius2Fahrenheit(int32(i))
		// fmt.Printf("Comparing %v to %v\n", int32(i), control)
		if tempF != control {
			return fmt.Errorf("%v got %v, expected %v", name, tempF, control)
		}
	}
	return nil
}
