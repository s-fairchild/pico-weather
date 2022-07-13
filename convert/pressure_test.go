package convert

import (
	"testing"
)

func TestMilliPa2MilliBar(t *testing.T) {


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
