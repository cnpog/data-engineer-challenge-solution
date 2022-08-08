package main

import (
	"context"
	"data-engineer-challenge/pkg/counting"
	"data-engineer-challenge/pkg/input"
	"data-engineer-challenge/pkg/kafkaadapter"
	"data-engineer-challenge/pkg/metric"
	"data-engineer-challenge/pkg/output"
	"data-engineer-challenge/pkg/stdout"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

// init function runs before main
// makes sure that the topics exist
func init() {
	kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKAIP"), os.Getenv("KAFKATOPIC"), 0)
}
func main() {
	log.Println("starting")
	// Create a new Kafka reader adapter.
	kafkaReaderAdapter := kafkaadapter.NewKafkaReaderAdapter(os.Getenv("KAFKAIP"), os.Getenv("KAFKATOPIC"))
	// initializes input service with the Kafka reader adapter.
	inputService := input.NewInputService(kafkaReaderAdapter)
	// Get the input channel from the input service.
	inputChan := inputService.GetInputChannel()
	// Stop the input service when the main function exits.
	defer inputService.StopInputChannel()

	// Create a new counting service with a 1 Minute sliding window interval
	countingService := counting.NewCountingService(time.Minute)
	// Get the output channel from the counting service.
	outputChan := countingService.GetOutputChannel()
	// Stop the counting service when the main function exits.
	defer countingService.StopOutputChannel()

	// init metric service
	metricService := metric.NewMetricService()
	// go function that reads from the input channel and counts the events.
	go func() {
		for event := range inputChan {
			metricService.AddFrame()
			countingService.Count(event)
		}
	}()
	// start logging fps
	metricService.LogFps()
	// Create a new console logger.
	console := stdout.NewConsoleLogger()
	outService := output.NewOutputService(console)
	// reads from the output channel and sends the events to the output service.
	for outgoingEvent := range outputChan {
		err := outService.Write(outgoingEvent)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
