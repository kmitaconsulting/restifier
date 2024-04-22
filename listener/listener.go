package listener

import (
	"kmitaconsulting/restifier/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Listener struct {
	consumer *kafka.Consumer
}

func NewListener() *Listener {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": data.KafkaServer, "group.id": data.KafkaGroup, "auto.offset.reset": "latest"})
	if err != nil {
		panic(err)
	}
	return &Listener{consumer: c}
}

func (l *Listener) Listen() {
	
}
