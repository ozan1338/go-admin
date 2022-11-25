package events

import (
	"encoding/json"
	"go-admin/src/database"
	"go-admin/src/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Listen(message *kafka.Message) error {
	key := string(message.Key)
	
	switch key {
	case "link_created":
		return LinkCreated(message.Value)
	}
	return nil
}

func LinkCreated(value []byte) error {
	var link models.Link

	if err := json.Unmarshal(value, &link); err != nil {
		return err
	}

	if err := database.DB.Create(&link).Error; err != nil {
		return err
	}

	return nil
}