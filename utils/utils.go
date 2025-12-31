package utils

import kfk "github.com/segmentio/kafka-go"

func ConnectToKafka() {
	kfk.Dial("tcp", "localhost:9092")
}
