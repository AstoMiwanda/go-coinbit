package producer

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// KafkaProducer hold kafka producer session
type KafkaProducer struct {
	Producer sarama.SyncProducer
}

// SendMessage function to send message into kafka
func (p *KafkaProducer) SendMessage(topic string, key string, data []byte) error {

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		logrus.Errorf("Send message error: %v", err)
		return err
	}

	logrus.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return nil
}
