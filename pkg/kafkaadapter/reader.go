package kafkaadapter

import (
	"context"
	"data-engineer-challenge/pkg/input"
	"encoding/json"
	"strings"

	"github.com/segmentio/kafka-go"
)

type KafkaReaderAdapter struct {
	reader *kafka.Reader
}

func NewKafkaReaderAdapter(kafkaURL, topic string) *KafkaReaderAdapter {
	brokers := strings.Split(kafkaURL, ",")
	return &KafkaReaderAdapter{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		}),
	}
}

func (k *KafkaReaderAdapter) Read() (input.Event, error) {
	m, err := k.reader.ReadMessage(context.Background())
	if err != nil {
		return input.Event{}, err
	}
	var event input.Event
	err = json.Unmarshal(m.Value, &event)
	if err != nil {
		return input.Event{}, err
	}
	return event, nil
}
