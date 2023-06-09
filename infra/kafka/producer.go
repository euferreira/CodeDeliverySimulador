package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		//"security.protocol": os.Getenv("security.protocol"),
		//"sasl.mechanisms": os.Getenv("sasl.mechanisms"),
		//"sasl.username": os.Getenv("sasl.username"),
		//"sasl.password": os.Getenv("sasl.password"),
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		//TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: -1},
		Value:          []byte(msg),
	}
	log.Println("Msg sending...");

	err := producer.Produce(message, nil)

	if err != nil {
		log.Println("DEU RUIM :(")
		return err
	}
	return nil
}
