package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"	
)

const (
	kafkaHost = "host.docker.internal:9092"
	topic     = "fancy-topic"
)

func main() {
	// Create a new Kafka producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Metadata.AllowAutoTopicCreation = false

	conn, err := sarama.NewSyncProducer([]string{kafkaHost}, config)
	if err != nil {
		log.Fatal("Failed to start Kafka producer: ", err)
	}

	defer conn.Close()

	i := 0
	for {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("Key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Fancy message %d", i)),
		}

		partition, offset, err := conn.SendMessage(msg)
		if err != nil {
			log.Println("Failed to send message: ", err)
		} else {
			log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		}

		i++
		time.Sleep(time.Second)
	}
}