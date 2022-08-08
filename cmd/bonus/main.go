package main

import (
	"data-engineer-challenge/pkg/counting"
	"data-engineer-challenge/pkg/input"
	"data-engineer-challenge/pkg/metric"
	"data-engineer-challenge/pkg/output"
	"data-engineer-challenge/pkg/stdin"
	"data-engineer-challenge/pkg/stdout"
	"log"
	"time"
)

func main() {
	// reads the input from stdin
	fileReader := stdin.NewFileReader("./payload.json")
	// initializes input service with the file reader
	inputService := input.NewInputService(fileReader)
	// Get the input channel from the input service.
	inputChan := inputService.GetInputChannel()
	// Stop the input service when the main function exits.
	defer inputService.StopInputChannel()

	// Create a new counting service with a 3 second sliding window interval
	countingService := counting.NewCountingService(time.Second * 3)
	// Get the output channel from the counting service.
	outputChan := countingService.GetOutputChannel()
	// Stop the counting service when the main function exits.
	defer countingService.StopOutputChannel()

	// init second counting service with different window size
	countingService2 := counting.NewCountingService(time.Second * 5)
	outputChan2 := countingService2.GetOutputChannel()
	defer countingService2.StopOutputChannel()

	// init metric service
	metricService := metric.NewMetricService()
	// go function that reads from the input channel and counts the events.
	go func() {
		for event := range inputChan {
			metricService.AddFrame()
			countingService.Count(event)
			countingService2.Count(event)
		}
	}()
	// start logging fps
	metricService.LogFps()
	// Create a new console logger.
	console := stdout.NewConsoleLogger()
	outService := output.NewOutputService(console)
	// reads from the output channel and sends the events to the output service.
	go func() {
		for outgoingEvent := range outputChan {
			err := outService.Write(outgoingEvent)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}()
	// reads from the second output channel and sends the events to the output service.
	for outgoingEvent := range outputChan2 {
		err := outService.Write(outgoingEvent)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
