package counting

import (
	"sync"
	"time"
)

type Service interface {
	Count(evt Event)
	GetOutputChannel() <-chan OutgoingEvent
	StopOutputChannel()
	GetWindowSlides() []struct {
		UserMap   map[string]interface{}
		FirstTime *time.Time
	}
}

type countingService struct {
	window       time.Duration
	mu           sync.Mutex
	windowSlides []struct {
		UserMap   map[string]interface{}
		FirstTime *time.Time
	}
	shutdownChannel chan interface{}
	ticker          *time.Ticker
	buffer          int
}

func NewCountingService(window time.Duration) Service {
	return &countingService{
		window:          window,
		ticker:          time.NewTicker(window + time.Second*5),
		shutdownChannel: make(chan interface{}),
		buffer:          2,
	}
}

// Count adds an incoming event to the correct sliding window or creates a new sliding window
func (s *countingService) Count(evt Event) {
	var evtCounted bool
	evtTime := time.Unix(evt.Ts, 0)
	for _, slide := range s.windowSlides {
		if slide.FirstTime == nil {
			s.mu.Lock()
			slide.FirstTime = &evtTime
			s.mu.Unlock()
		}
		if evtTime.Unix()-slide.FirstTime.Unix() <= int64(s.window) && evtTime.Unix()-slide.FirstTime.Unix() > 0 {
			s.mu.Lock()
			slide.UserMap[evt.Uid] = struct{}{}
			s.mu.Unlock()
			evtCounted = true
			break
		}
	}
	if !evtCounted {
		userMap := make(map[string]interface{})
		userMap[evt.Uid] = struct{}{}
		firstTime := &evtTime
		s.mu.Lock()
		s.windowSlides = append(s.windowSlides, struct {
			UserMap   map[string]interface{}
			FirstTime *time.Time
		}{
			UserMap:   userMap,
			FirstTime: firstTime,
		})
		s.mu.Unlock()
	}
}

// GetOutputChannel returns a channel that will emit the number of unique users in the sliding window
func (s *countingService) GetOutputChannel() <-chan OutgoingEvent {
	ch := make(chan OutgoingEvent, s.buffer)
	go func() {
		for {
			select {
			case <-s.shutdownChannel:
				close(ch)
				return
			case <-s.ticker.C:
				if s.getUniqueUsersCount() > 0 {
					ch <- OutgoingEvent{
						UniqueUsers:       s.getUniqueUsersCount(),
						DurationInSeconds: int(s.window.Seconds()),
						Timestamp:         s.windowSlides[0].FirstTime.Unix(),
					}
					s.mu.Lock()
					s.windowSlides = append(s.windowSlides[:0], s.windowSlides[1:]...)
					s.mu.Unlock()
				}
			default:
				time.Sleep(time.Microsecond * 10)
			}

		}
	}()
	return ch
}

// StopOutputChannel stops the output channel
func (s *countingService) StopOutputChannel() {
	close(s.shutdownChannel)
}

// getUniqueUsersCount returns the number of unique users in the sliding window
func (s *countingService) getUniqueUsersCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.windowSlides) > 0 {
		return len(s.windowSlides[0].UserMap)
	}
	return 0
}

func (s *countingService) GetWindowSlides() []struct {
	UserMap   map[string]interface{}
	FirstTime *time.Time
} {
	return s.windowSlides
}
