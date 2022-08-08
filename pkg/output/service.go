package output

import "data-engineer-challenge/pkg/counting"

type Service interface {
	Write(outgoingEvent counting.OutgoingEvent) error
}

type OutputAdapter interface {
	Write(outputEvent OutputEvent) error
}
type outputService struct {
	adapter OutputAdapter
}

func NewOutputService(adapter OutputAdapter) Service {
	return &outputService{
		adapter: adapter,
	}
}

func (s *outputService) Write(outgoingEvent counting.OutgoingEvent) error {
	return s.adapter.Write(OutputEvent{
		Timestamp:         outgoingEvent.Timestamp,
		DurationInSeconds: outgoingEvent.DurationInSeconds,
		UniqueUsers:       outgoingEvent.UniqueUsers,
	})
}
