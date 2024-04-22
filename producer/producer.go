package producer

import (
	"encoding/json"
	"kmitaconsulting/restifier/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer() *Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": data.KafkaServer})
	if err != nil {
		panic(err)
	}
	return &Producer{producer: p}
}

func (p *Producer) ProduceSomeObject(someObject *data.SomeObject) {
	value, err := json.Marshal(someObject)
	if err != nil {
		panic(err)
	}

	topic := data.KafkaTopic

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	if err != nil {
		panic(err)
	}
}
