package main

import "github.com/Shopify/sarama"

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // leader和follower都需确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机partition
	config.Producer.Return.Successes = true                   // 确认 成功交付的消息将在success channel返回
	//连接kafka
	conn, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}

	topic, msg := "topic", "hello world"
	producerMessage := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(msg)}
	_, _, _ = conn.SendMessage(producerMessage)
}
