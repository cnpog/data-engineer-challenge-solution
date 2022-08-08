package counting_test

import (
	"data-engineer-challenge/pkg/counting"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test counting with all unique values
func TestCounting(t *testing.T) {
	countingService := counting.NewCountingService(time.Second * 5)

	for i := 0; i < 60; i++ {
		time.Sleep(time.Millisecond * 200)
		countingService.Count(counting.Event{
			Ts:  time.Now().Unix(),
			Uid: uuid.New().String(),
		})
	}
	assert.Greater(t, len(countingService.GetWindowSlides()[0].UserMap), 10)
	assert.Less(t, len(countingService.GetWindowSlides()[0].UserMap), 35)
	assert.Equal(t, 3, len(countingService.GetWindowSlides()))
}

// Test counting with duplicates
func TestCountingNotUnique(t *testing.T) {
	countingService := counting.NewCountingService(time.Second * 5)

	for i := 0; i < 60; i++ {
		time.Sleep(time.Millisecond)
		uid := ""
		if i%2 == 0 {
			uid = "1"
		} else {
			uid = uuid.New().String()
		}
		countingService.Count(counting.Event{
			Ts:  time.Now().Unix(),
			Uid: uid,
		})
	}
	assert.Equal(t, 31, len(countingService.GetWindowSlides()[0].UserMap))
	assert.Equal(t, 1, len(countingService.GetWindowSlides()))
}

func TestXxx(t *testing.T) {
	firsttime := time.Now()
	secondtime := firsttime.Add(time.Second * 1)
	if secondtime.Unix()-firsttime.Unix() <= int64(time.Second*5) && secondtime.Unix()-firsttime.Unix() > 0 {
		// assert.Equal(t, time.Now(), firsttime)
		// assert.Equal(t, secondtime.Unix()-firsttime.Unix(), int64(5))
	} else {
		t.Fail()
	}
}
