package windvane

import (
	"fmt"
	m "machine"
	"math"
)

const (
	R1  float64 = 10_000.0 // resistor 1 ohms
	vin float64 = 3.3 // 3.3 volts
)

var (
	vaneResistances = []float64{
		33_000,
		6.570,
		8.200,
		891,
		1000,
		688,
		2.200,
		1.410,
		3.900,
		3.140,
		1600,
		14.120,
		120_000,
		42.120,
		64.900,
		21.880,
	}
	possibleVolts = map[float64]float32{}
)

func GetWindDegrees(wv m.ADC) (float64, error) {
	rawReading := wv.Get()
	// TODO fix voltage calculation
	// create function to calculate voltage here
	reading := float32((float32(rawReading) / 10_000_00.0) * 100.0)
	for d, pV := range possibleVolts {
		if float32(reading) / 1000 == pV {
			// return degree of wind
			return math.Round(float64(possibleVolts[d])), nil
		}
	}
	return 999, fmt.Errorf("Unable to detect wind degree, voltage reading was: %f", reading)
}

func InitWindvane(r1 float64) m.ADC {
	m.InitADC()
	
	possibleVolts = generatePossibleVoltages(r1)
	windVane := m.ADC{Pin: m.ADC0}
	windVane.Configure(m.ADCConfig{})
	return windVane
}

func generatePossibleVoltages(r1 float64) map[float64]float32 {

	var degree float32 = 0
	pV := make(map[float64]float32, len(vaneResistances))
	for _, r2 := range vaneResistances {
		vout := voltageDivider(r2, r1)
		pV[vout] = degree
		// degrees in between each resistor in windvane
		degree += 22.5
		fmt.Printf("Possible Voltage: %f\nAt Degree %f\n", vout, degree)
	}
	return pV
}

func voltageDivider(r2, r1 float64) float64 {
	return ((vin * r2) / (r1 + r2) * 100.0)
	// return (vin * r2) / (r1 + r2)
}
