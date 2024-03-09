package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"lee-activity-framework/lee-activity-api/mq/kafka"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
)

func InitKafka() {
	brokers := strings.Split("localhost:9092", ",")
	// 创建消费者
	consumer, err := createConsumer(brokers, "lee-activity-task")
	if err != nil {
		log.Fatal("创建消费者失败:", err)
	}
	defer func() {
		if err = consumer.Close(); err != nil {
			log.Fatal("关闭消费者失败:", err)
		}
	}()

	topics := []string{new(kafka.SendGiftMsg).Topic(), new(kafka.SendMessageMsg).Topic()}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		consumeMessages(consumer, topics)
	}()

	// 监听退出信号
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm

	// 优雅关闭消费者
	wg.Wait()
}

func createConsumer(brokers []string, groupId string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	return sarama.NewConsumerGroup(brokers, groupId, config)
}

type KafkaConsumerGroupHandler struct {
	ready chan bool
}

func (handler *KafkaConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	close(handler.ready)
	return nil
}

func (handler *KafkaConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (handler *KafkaConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("消息: 主题=%s 分区=%d 偏移量=%d\n", message.Topic, message.Partition, message.Offset)
		fmt.Printf("消息内容: %s\n", string(message.Value))
		sess.MarkMessage(message, "")
		// TODO：业务逻辑
	}
	return nil
}

func consumeMessages(consumer sarama.ConsumerGroup, topics []string) {
	handler := &KafkaConsumerGroupHandler{
		ready: make(chan bool),
	}

	for {
		err := consumer.Consume(context.Background(), topics, handler)
		if err != nil {
			fmt.Printf("消费者错误:%v", err)
		}

		select {
		case <-handler.ready:
		default:
			return
		}
	}
}
