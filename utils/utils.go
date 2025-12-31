package utils

import (
	"log"

	kfk "github.com/segmentio/kafka-go"
)

func ConnectToKafka() {
	_, err := kfk.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Println(err)
	}

}
