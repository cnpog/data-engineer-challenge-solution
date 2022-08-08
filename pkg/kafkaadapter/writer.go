package kafkaadapter

import (
	"context"
	"data-engineer-challenge/pkg/output"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type KafkaWriterAdapter struct {
	writer *kafka.Writer
}

func NewKafkaWriterAdapter(kafkaURL, topic string) *KafkaWriterAdapter {
	return &KafkaWriterAdapter{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(kafkaURL),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (k *KafkaWriterAdapter) Write(outEvent output.OutputEvent) error {
	output, _ := json.Marshal(outEvent)
	msg := kafka.Message{
		Value: output,
	}
	return k.writer.WriteMessages(context.Background(), msg)
}
