package kafka

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
)

const legalEntitiesTopic = "legal-entities-created"

type LegalEntitySender struct {
	producer sarama.SyncProducer
}

func NewLegalEntitySender() *LegalEntitySender {
	return &LegalEntitySender{producer: newSyncProducer()}
}

type LegalEntityCreatedMessage struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// SendLegalEntityCreated публикует событие; в degraded-mode молча пропускает.
func (s *LegalEntitySender) SendLegalEntityCreated(id string, createdAt time.Time) error {
	if s.producer == nil {
		return nil
	}
	payload, err := json.Marshal(LegalEntityCreatedMessage{
		ID:        id,
		CreatedAt: createdAt.UTC(),
	})
	if err != nil {
		return err
	}
	_, _, err = s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: legalEntitiesTopic,
		Value: sarama.ByteEncoder(payload),
	})
	return err
}
