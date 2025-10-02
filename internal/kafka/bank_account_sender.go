package kafka

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
)

const bankAccountsTopic = "bank-accounts-created"

type BankAccountSender struct {
	producer sarama.SyncProducer
}

func NewBankAccountSender() *BankAccountSender {
	return &BankAccountSender{producer: newSyncProducer()}
}

type BankAccountCreatedMessage struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *BankAccountSender) SendBankAccountCreated(id string, createdAt time.Time) error {
	if s.producer == nil {
		return nil
	}
	payload, err := json.Marshal(BankAccountCreatedMessage{
		ID:        id,
		CreatedAt: createdAt.UTC(),
	})
	if err != nil {
		return err
	}
	_, _, err = s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: bankAccountsTopic,
		Value: sarama.ByteEncoder(payload),
	})
	return err
}
