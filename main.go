package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/leviatan89/temperatureMonitor/reader"
	"github.com/leviatan89/temperatureMonitor/sendtoapi"
)

func main() {
	err := godotenv.Load(".env") // Get the environment variables from a file called .env
	if err != nil {
		log.Println(err)
	}
	for _ = range time.Tick(30 * time.Second) {
		read := reader.ReadNow()
		if read.Valid {
			sendtoapi.Sendtoapi(read)
		}
	}
}
