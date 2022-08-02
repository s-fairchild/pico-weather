package bucket

import (
	"testing"
)

func TestCalculateRain(t *testing.T) {

	bucketSizemm := 0.2794
	expect := 2.2352
	n := calculateRain(8, bucketSizemm)
	if n != expect {
		t.Fatalf("%v expected %v got %v", t.Name(), expect, n)
	}

	bucketSizeInches := 0.011
	expect = 0.088
	n = calculateRain(8, bucketSizeInches)
	if n != expect {
		t.Fatalf("%v expected %v got %v", t.Name(), expect, n)
	}
}
