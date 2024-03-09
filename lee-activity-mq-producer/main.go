package main

import (
	"lee-activity-framework/lee-activity-mq-producer/kafka"
	"lee-activity-framework/lee-activity-mq-producer/server/http"
)

func main() {
	// Kafka
	kafka.InitKafka()
	// Http
	http.InitRouter()
}
