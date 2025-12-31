package helper

import "fmt"

const KAFKABROKER = "broke:9092"

func ConnectToBroker() {
	fmt.Println(KAFKABROKER)
	fmt.Println("go lib test v1")
}
