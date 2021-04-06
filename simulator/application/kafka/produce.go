package kafka

import (
	"encoding/json"
	"github.com/gabitchian/simulator-fsfc/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route2 "github.com/gabitchian/simulator-fsfc/application/route"
	"log"
	"os"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()

	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}