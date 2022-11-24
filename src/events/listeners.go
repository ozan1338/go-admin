package events

import (
	"encoding/json"
	"go-admin/src/database"
	"go-admin/src/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Listen(message *kafka.Message) {
	key := string(message.Key)
	
	switch key {
	case "link_created":
		LinkCreated(message.Value)
	}
}

func LinkCreated(value []byte) {
	var link models.Link

	json.Unmarshal(value, &link)

	database.DB.Create(&link)
}