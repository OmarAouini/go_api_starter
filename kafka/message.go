package kafka

import (
	"time"

	"github.com/google/uuid"
)

type KafkaMessage struct {
	UUID      uuid.UUID   `json:"uuid"`
	Timestamp time.Time   `json:"timestamp"`
	Content   interface{} `json:"content"`
}
