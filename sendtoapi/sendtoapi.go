package sendtoapi

import (
	"fmt"
	"log"
	"os"

	adafruitio "github.com/adafruit/io-client-go"
	"github.com/leviatan89/temperatureMonitor/reader"
)

func Sendtoapi(read reader.Read) {

	client := adafruitio.NewClient(os.Getenv("ADAFRUIT_IO_KEY"))

	// HUMIDITY
	feed, _, err := client.Feed.Get("room-humidity")
	if err != nil {
		log.Printf("unable to load Feed with key %v")
		return
	}

	client.SetFeed(feed)
	val := &adafruitio.Data{Value: fmt.Sprintf("%f", read.Humidity)}

	_, _, err = client.Data.Send(val) // The response is: data, rsponse, err
	if err != nil {
		log.Println("--- json unmarshal error? We can ignore this error as it looks like there is something broken in the library ---")
		log.Println(err)
	}

	// TEMPERATURE
	feed, _, err = client.Feed.Get("room-temperature")
	if err != nil {
		log.Println(err)
		return
	}
	client.SetFeed(feed)
	val = &adafruitio.Data{Value: fmt.Sprintf("%f", read.Temperature)}
	_, _, err = client.Data.Send(val) // The response is: data, response, err
	if err != nil {
		log.Println("--- json unmarshal error? We can ignore this error as it looks like there is something broken in the library ---")
		log.Println(err)
	}

	return
}
