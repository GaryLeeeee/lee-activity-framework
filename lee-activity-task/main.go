package main

import (
	"lee-activity-framework/lee-activity-task/dao"
	"lee-activity-framework/lee-activity-task/kafka"
	"lee-activity-framework/lee-activity-task/server/http"
)

func main() {
	// Kafka
	go kafka.InitKafka()
	// MySQL
	dao.InitMySQL()
	// Redis
	dao.InitRedis()
	// Http
	http.InitRouter()
}
