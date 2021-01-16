package reader

import (
	"fmt"
	"time"

	"github.com/MichaelS11/go-dht"
)

type Read struct {
	Humidity    float64
	Temperature float64
	Valid       bool
	Timestamp   time.Time
}

func ReadNow() Read {
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return Read{}
	}

	dht, err := dht.NewDHT("GPIO17", dht.Celsius, "")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return Read{}
	}
	reading := Read{Valid: true}
	reading.Humidity, reading.Temperature, err = dht.ReadRetry(30)
	if err != nil {
		fmt.Println("Read error:", err)
		return Read{Valid: false}
	}
	reading.Timestamp = time.Now()
	return reading
}
