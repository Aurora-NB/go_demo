package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	servers := []string{"47.108.93.159:9092"}
	topic := "log"

	consumer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		return
	}
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return
	}
	for _, partition := range partitions {
		consumePartition, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		if err != nil {
			return
		}
		defer consumePartition.AsyncClose()
		go func() {
			for m := range consumePartition.Messages() {
				fmt.Println(string(m.Value))
			}
		}()
	}
	select {}
}
