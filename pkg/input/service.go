package input

import (
	"data-engineer-challenge/pkg/counting"
)

type InputService interface {
	GetInputChannel() <-chan counting.Event
	StopInputChannel()
}

type InputAdapter interface {
	Read() (Event, error)
}
type inputService struct {
	adapter         InputAdapter
	buffer          int
	shutdownChannel chan interface{}
}

func NewInputService(adapter InputAdapter) InputService {
	return &inputService{
		adapter:         adapter,
		buffer:          2,
		shutdownChannel: make(chan interface{}),
	}
}

// GetInputChannel returns a channel that will be populated with events.
func (s *inputService) GetInputChannel() <-chan counting.Event {
	ch := make(chan counting.Event, s.buffer)
	go func() {
		for {
			select {
			case <-s.shutdownChannel:
				close(ch)
				return
			default:
				event, err := s.adapter.Read()
				if err != nil {
					continue
				}
				ch <- counting.Event{
					Ts:  event.Ts,
					Uid: event.Uid,
				}
			}

		}
	}()
	return ch
}

// StopInputChannel stops the input channel.
func (s *inputService) StopInputChannel() {
	close(s.shutdownChannel)
}
