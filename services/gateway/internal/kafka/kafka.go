package kafka

import (
	"Chat/pkg/models"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/dchest/uniuri"
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

// Соединение с брокером
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

// Отправление сообщения в брокер по топику
func (kfk *KafkaConnection) SendMessage(msg models.Message, topic string) error {
	log.Println("Writing message:", msg)
	conn := kfk.ProducerTopics[topic]
	msg.Hash = uniuri.New()
	jsonedMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	writtenBytes, err := conn.Write(jsonedMsg)
	log.Println("Bytes written:", writtenBytes)
	if err != nil {
		return err
	}
	return nil
}

// Получение сообщений из брокера по топику
func (kfk *KafkaConnection) OpenMessageTube(ch *chan models.BeautifiedMessage, topic string) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kfk.ConsumerTopics[topic].RemoteAddr().String()},
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6,
	})
	lastMessageHash := ""
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Cant read message:", err)
		}
		log.Println("Got message from broker:", string(msg.Value))
		readMessage := models.BeautifiedMessage{}
		err = json.Unmarshal(msg.Value, &readMessage)
		if err != nil {
			log.Println("Cant unmarshal message:", err)
		}
		if readMessage.Hash != lastMessageHash {
			log.Println("Sending message to gateway")
			lastMessageHash = readMessage.Hash
			*ch <- readMessage
		} else {
			log.Println("Already got message with the same hash")
		}

	}
}
