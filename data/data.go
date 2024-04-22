package data

const (
	KafkaServer = "localhost:9092"
	KafkaTopic = "some.topic"
	KafkaGroup = "example_group"
)

type SomeObject struct {
	ID string
	SomeData string
	Status bool
}
