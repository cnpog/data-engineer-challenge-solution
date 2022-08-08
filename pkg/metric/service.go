package metric

import (
	"log"
	"time"
)

type Service interface {
	AddFrame()
	LogFps()
}

type service struct {
	frameCount int
	ticker     *time.Ticker
}

func NewMetricService() Service {
	return &service{
		frameCount: 0,
		ticker:     time.NewTicker(time.Second),
	}
}

func (s *service) AddFrame() {
	s.frameCount++
}

// LogFps logs the fps every second
// ideally you would not directly log into the console here
func (s *service) LogFps() {
	go func() {
		for {
			select {
			case <-s.ticker.C:
				if s.frameCount > 0 {
					log.Printf("fps: %d", s.frameCount)
					s.frameCount = 0
				}
			default:
				// prevent goroutine to take too much cpu time
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
}
