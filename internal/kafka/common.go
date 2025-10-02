package kafka

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

const (
	defaultBrokers = "localhost:29092"
)

func brokersFromEnv() []string {
	raw := strings.TrimSpace(os.Getenv("KAFKA_BROKERS"))
	if raw == "" {
		return strings.Split(defaultBrokers, ",")
	}
	return strings.Split(raw, ",")
}

// newSyncProducer создает idempotent SyncProducer.
// При ошибке возвращает nil и логирует предупреждение (degraded mode).
func newSyncProducer() sarama.SyncProducer {
	cfg := sarama.NewConfig()

	// Idempotent producer + гарантии доставки
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Idempotent = true
	cfg.Net.MaxOpenRequests = 1
	cfg.Producer.Retry.Max = 10
	cfg.Producer.Retry.Backoff = 200 * time.Millisecond

	// Версия кластера (совместимо с CP 7.3, Kafka 3.4.x)
	cfg.Version = sarama.V3_4_0_0

	prod, err := sarama.NewSyncProducer(brokersFromEnv(), cfg)
	if err != nil {
		log.Printf("kafka: producer init failed: %v (degraded mode; events will be skipped)", err)
		return nil
	}
	return prod
}
