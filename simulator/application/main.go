package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafka2 "github.com/kaualimamartins/simulator/application/kafka"
	"github.com/kaualimamartins/simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)

	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))

		go kafka2.Produce(msg)
	}

	// router := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// router.LoadPositions()
	// stringJson, _ := router.ExportPositionsToJson()

	// fmt.Println(stringJson[1])
}
