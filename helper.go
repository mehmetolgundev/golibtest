package helper

import (
	"fmt"

	"github.com/mehmetolgundev/golibtest/utils"
)

const KAFKABROKER = "broke:9092"

func ConnectToBroker() {
	fmt.Println(KAFKABROKER)
	fmt.Println("go lib test v2")
	utils.ConnectToKafka()
}
