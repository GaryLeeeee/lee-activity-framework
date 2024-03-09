package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"strings"
	"time"
)

var _producer sarama.AsyncProducer

func InitKafka() {
	brokers := strings.Split("localhost:9092", ",")
	// 创建生产者
	producer, err := createProducer(brokers)
	if err != nil {
		log.Fatal("创建生产者失败:", err)
	}
	_producer = producer
	// TODO：producer close
	//defer func() {
	//	if err = producer.Close(); err != nil {
	//		log.Fatal("关闭生产者失败:", err)
	//	}
	//}()
}

func createProducer(brokers []string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	return sarama.NewAsyncProducer(brokers, config)
}

func ProduceMessage(topic, value string) {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	}
	_producer.Input() <- message
}
