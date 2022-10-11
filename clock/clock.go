package clock

import (
	m "machine"
	"time"
	"fmt"

	"tinygo.org/x/drivers/ds1307"
	"tinygo.org/x/drivers/i2csoft"
)

func SetTime(rtc *ds1307.Device) {

	rtc.SetTime(time.Date(2022, 10, 7, 18, 15, 12, 0, time.UTC))
}

// InitRtc creates a software i2c interface using sda and scl using i2csoft package
func InitRtc(sda, scl m.Pin) (ds1307.Device, error) {
	i2c := i2csoft.New(scl, sda)
	err := i2c.Configure(i2csoft.I2CConfig{})
	if err != nil {
		return ds1307.Device{}, err
	}

	return ds1307.New(i2c), nil
}

func FormatTime(rtc *ds1307.Device) (string, error) {

	rt, err := rtc.ReadTime()
	if err != nil {
		return "", err
	}

	// println(t.Hour(), ":", t.Minute(), ":", t.Second(), " ", t.Month(), "/", t.Day(), "/", t.Year())
	return fmt.Sprintf("%v:%v:%v %v/%d/%v", rt.Hour(), rt.Minute(), rt.Second(), rt.Day(), rt.Month(), rt.Year()), nil
}

func Elapsed(t time.Time) string {

	hour, min, sec := t.Clock()
	return fmt.Sprintf("Time elapsed: %v:%v:%v", hour, min, sec)
}

func ExecutionTime(then time.Time) string {

	return fmt.Sprintf("Calculation time:  %v", time.Since(then))
}
