package dao

import (
	"fmt"
)

const (
	KafkaUUIDKey    = "lee-activity:task:kafka_uuid:%s" // kafka uuid幂等key
	KafkaUUIDExpire = 86400
)

// Kafka幂等
func CheckKafkaUUIDExists(uuid string) bool {
	ret, err := _rd.Do("SET", fmt.Sprintf(KafkaUUIDKey, uuid), "1", "EX", KafkaUUIDExpire, "NX").Result()
	if err == nil && ret == "OK" {
		return false
	}
	return true
}
