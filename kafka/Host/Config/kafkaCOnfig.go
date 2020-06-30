package Config

import (
	"github.com/Shopify/sarama"
	"time"
)

func GetKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()

	kafkaConfig.Producer.Return.Successes = true

	kafkaConfig.Net.WriteTimeout = 5 * time.Second

	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {

		kafkaConfig.Net.SASL.Enable = true

		kafkaConfig.Net.SASL.User = username

		kafkaConfig.Net.SASL.Password = password

	}

	return kafkaConfig
}