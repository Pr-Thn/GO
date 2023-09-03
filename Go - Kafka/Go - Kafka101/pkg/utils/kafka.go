package utils

import (
	"context"
	"log"

	"kafka/config"

	"github.com/segmentio/kafka-go"
)

// When Connection between apps and kafka
// Import Kafka config Struct and Export struct
func KafkaConn(cfg config.KafkaConnCfg) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Url, cfg.Topic, 0)
	if err != nil {
		log.Fatal("Failed to Dialleader:", err)
	}
	return conn
}

// Check if their topic is empty or not if not create one
// return bool true or if it have or haven't
func IsTopicAlreadyExists(conn *kafka.Conn, topic string) bool {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range partitions {
		if p.Topic == topic {
			return true
		}
	}
	return false
}
