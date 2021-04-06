package main

import (
	"fmt"
	//route2 "github.com/gabitchian/simulator-fsfc/application/route"
	kafka2 "github.com/gabitchian/simulator-fsfc/application/kafka"
	"github.com/gabitchian/simulator-fsfc/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
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
	//route := route2.Route{
	//	ID: "1",
	//	ClientID: "1",
	//}
	//
	//route.LoadPositions()
	//stringJson, _ := route.ExportJsonPositions()
	//fmt.Println(stringJson[0])
}
