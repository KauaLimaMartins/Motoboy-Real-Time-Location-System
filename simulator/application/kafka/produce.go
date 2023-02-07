package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	router "github.com/kaualimamartins/simulator/application/route"
	"github.com/kaualimamartins/simulator/infra/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := router.NewRoute()

	json.Unmarshal(msg.Value, &route)

	route.LoadPositions()

	positions, err := route.ExportPositionsToJson()
	if err != nil {
		log.Println(err)
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
