package kafka

import (
	"Chat/pkg/models"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/segmentio/kafka-go"
)

type KafkaConnection struct {
	ProducerTopics map[string]*kafka.Conn
	ConsumerTopics map[string]*kafka.Conn
}

type Config struct {
	Host           string `yaml:"kafka_host"`
	Port           string `yaml:"kafka_port"`
	ProducerTopics string `yaml:"producer_topics"`
	ConsumerTopics string `yaml:"consumer_topics"`
}

var partition = 0

func New(cfg *Config) (*KafkaConnection, error) {
	kafkaConnection := KafkaConnection{ProducerTopics: map[string]*kafka.Conn{}, ConsumerTopics: map[string]*kafka.Conn{}}
	producerTopics := strings.Split(cfg.ProducerTopics, " ")
	log.Println("connecting to broker on " + cfg.Host + cfg.Port)
	for _, elem := range producerTopics {
		conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Host+cfg.Port, elem, partition)
		if err != nil {
			return nil, err
		}
		kafkaConnection.ProducerTopics[elem] = conn
	}
	consumerTopics := strings.Split(cfg.ConsumerTopics, " ")
	for _, elem := range consumerTopics {
		conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Host+cfg.Port, elem, partition)
		if err != nil {
			return nil, err
		}
		kafkaConnection.ConsumerTopics[elem] = conn
	}
	return &kafkaConnection, nil
}

func (kfk *KafkaConnection) SendMessage(msg models.BeautifiedMessage, topic string) error {
	log.Println("Sending message:", msg)
	conn := kfk.ProducerTopics[topic]
	jsonedMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(jsonedMsg)
	if err != nil {
		return err
	}
	return nil
}

func (kfk *KafkaConnection) OpenMessageTube(ch *chan models.Message, topic string) error {
	conn := kfk.ConsumerTopics[topic]
	log.Println("Connection address:", conn.RemoteAddr())
	batch := conn.ReadBatch(70, 1e6)
	b := make([]byte, 10)
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		log.Println("Got message: ", b[:n])
		msg := models.Message{}
		err = json.Unmarshal(b[:n], &msg)
		if err != nil {
			break
		}
		*ch <- msg
	}
	if err := batch.Close(); err != nil {
		return err
	}
	if err := conn.Close(); err != nil {
		return err
	}
	return nil
}
